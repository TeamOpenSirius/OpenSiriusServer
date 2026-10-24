package masterdata

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"

	"ymst_server/apiwire"
	"ymst_server/core/config"
)

var backendHTTPClient = &http.Client{Timeout: 60 * time.Second}

type BackendManifest struct {
	URI              string
	SasToken         string
	Version          string
	PublishTimestamp int64
}

func SyncFromBackend(ctx context.Context, game config.GameConfig, log *zap.Logger) (config.GameConfig, error) {
	base := strings.TrimRight(strings.TrimSpace(game.MasterDataBackendURL), "/")
	if base == "" {
		return game, nil
	}
	manifest, err := fetchBackendManifest(ctx, base)
	if err != nil {
		return game, err
	}
	if strings.TrimSpace(manifest.URI) == "" {
		return game, errors.New("backend master manifest uri is empty")
	}
	root := strings.TrimSpace(game.MasterDataRoot)
	if root == "" {
		root = "data/master-data/production"
	}
	localPath, err := safeJoin(root, manifest.URI)
	if err != nil {
		return game, err
	}
	if manifest.Version != "" && manifest.Version == game.MasterDataVersion && fileExists(localPath) {
		if log != nil {
			log.Info("master data already synced",
				zap.String("version", manifest.Version),
				zap.String("path", localPath),
			)
		}
		return applyBackendManifest(game, manifest, root), nil
	}
	url := base + "/master-data/production/" + strings.TrimLeft(strings.ReplaceAll(manifest.URI, "\\", "/"), "/")
	if manifest.SasToken != "" {
		url += manifest.SasToken
	}
	if err := downloadMasterFile(ctx, url, localPath, log); err != nil {
		return game, err
	}
	return applyBackendManifest(game, manifest, root), nil
}

func fetchBackendManifest(ctx context.Context, base string) (BackendManifest, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+"/api/data/master", nil)
	if err != nil {
		return BackendManifest{}, err
	}
	req.Header.Set("Accept", apiwire.MsgpackContentType)
	req.Header.Set("User-Agent", "BestHTTP/2 v2.8.5")
	resp, err := backendHTTPClient.Do(req)
	if err != nil {
		return BackendManifest{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendManifest{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return BackendManifest{}, fmt.Errorf("backend master manifest http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	values, err := decodeMsgpackStream(body)
	if err != nil {
		return BackendManifest{}, err
	}
	env := normalizeEnvelope(values)
	if len(env) < 2 {
		return BackendManifest{}, fmt.Errorf("backend master manifest envelope too short: %d", len(env))
	}
	main, ok := env[1].([]any)
	if !ok {
		return BackendManifest{}, fmt.Errorf("backend master manifest main has type %T", env[1])
	}
	return BackendManifest{
		URI:              stringAt(main, 0),
		SasToken:         stringAt(main, 1),
		Version:          stringAt(main, 2),
		PublishTimestamp: int64At(main, 3),
	}, nil
}

func decodeMsgpackStream(body []byte) ([]any, error) {
	dec := msgpack.NewDecoder(bytes.NewReader(body))
	var values []any
	for {
		var v any
		err := dec.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		values = append(values, apiwire.NormalizeMsgpackValue(v))
	}
	return values, nil
}

func normalizeEnvelope(values []any) []any {
	if len(values) == 1 {
		if arr, ok := values[0].([]any); ok {
			return arr
		}
	}
	return values
}

func downloadMasterFile(ctx context.Context, url string, path string, log *zap.Logger) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "BestHTTP/2 v2.8.5")
	resp, err := backendHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("backend master file http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	tmp := path + ".part"
	out, err := os.Create(tmp)
	if err != nil {
		return err
	}
	var written int64
	nextLogAt := int64(1 << 20)
	buf := make([]byte, 256*1024)
	for {
		n, readErr := resp.Body.Read(buf)
		if n > 0 {
			if _, err := out.Write(buf[:n]); err != nil {
				_ = out.Close()
				return err
			}
			written += int64(n)
			if log != nil && written >= nextLogAt {
				log.Info("master data download progress",
					zap.Int64("downloaded", written),
					zap.Int64("total", resp.ContentLength),
					zap.String("path", path),
				)
				nextLogAt = written + int64(1<<20)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			_ = out.Close()
			return readErr
		}
	}
	if err := out.Close(); err != nil {
		return err
	}
	if resp.ContentLength > 0 && written != resp.ContentLength {
		return fmt.Errorf("master data size mismatch: got %d want %d", written, resp.ContentLength)
	}
	_ = os.Remove(path)
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	if log != nil {
		log.Info("master data downloaded",
			zap.Int64("size", written),
			zap.String("path", path),
		)
	}
	return nil
}

func applyBackendManifest(game config.GameConfig, manifest BackendManifest, root string) config.GameConfig {
	game.MasterDataRoot = root
	game.MasterDataURI = strings.TrimLeft(strings.ReplaceAll(manifest.URI, "\\", "/"), "/")
	if manifest.Version != "" {
		game.MasterDataVersion = manifest.Version
	}
	if manifest.PublishTimestamp > 0 {
		game.MasterDataPublishTimestamp = manifest.PublishTimestamp
	}
	return game
}

func safeJoin(root string, rel string) (string, error) {
	rel = strings.TrimLeft(strings.ReplaceAll(rel, "\\", "/"), "/")
	if rel == "" || strings.Contains(rel, "..") {
		return "", fmt.Errorf("invalid relative master path %q", rel)
	}
	return filepath.Join(root, filepath.FromSlash(rel)), nil
}

func fileExists(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir() && st.Size() > 0
}

func stringAt(arr []any, idx int) string {
	if idx < 0 || idx >= len(arr) {
		return ""
	}
	if s, ok := arr[idx].(string); ok {
		return s
	}
	return ""
}

func int64At(arr []any, idx int) int64 {
	if idx < 0 || idx >= len(arr) {
		return 0
	}
	return apiwire.Int64Value(arr[idx])
}
