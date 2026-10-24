package core

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/core/logger"
	"ymst_server/masterdata"
	"ymst_server/utils/fsutil"
)

const appDataDir = "data"

type AppCommon struct {
	Logger   *logger.Logger
	Config   config.Config
	Security config.SecurityConfig
	Database *database.Client
	GinMode  string

	cleanup func()
}

var Common *AppCommon

func InitCommon() error {
	if err := checkAppDir(); err != nil {
		return fmt.Errorf("init data directory failed: %w", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load failed: %w", err)
	}
	securityCfg, err := config.LoadSecurity()
	if err != nil {
		return fmt.Errorf("security config load failed: %w", err)
	}

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	logr, cleanup, err := logger.New(cfg.LogPath, cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("logger init failed: %w", err)
	}
	logger.RedirectGinLogs(logr)

	dbBootstrap, err := database.Bootstrap(gin.Mode(), logr)
	if err != nil {
		cleanup()
		return fmt.Errorf("database bootstrap failed: %w", err)
	}
	if _, err := masterdata.ImportDecodedJSON(context.Background(), dbBootstrap.Client.DB(), cfg.Game.MasterDecodedJSONRoot, logr.GetZap()); err != nil {
		_ = dbBootstrap.Client.Close()
		cleanup()
		return fmt.Errorf("master decoded json import failed: %w", err)
	}

	Common = &AppCommon{
		Logger:   logr,
		Config:   cfg,
		Security: securityCfg,
		Database: dbBootstrap.Client,
		GinMode:  gin.Mode(),
		cleanup:  cleanup,
	}
	return nil
}

func ShutdownCommon() {
	if Common == nil {
		return
	}
	if Common.Database != nil {
		if err := Common.Database.Close(); err != nil {
			Common.Logger.GetZap().Warn("database close failed", zap.Error(err))
		}
	}
	if Common.cleanup != nil {
		Common.cleanup()
	}
}

func checkAppDir() error {
	return fsutil.EnsureDir(appDataDir)
}

func GetSubDir(name string) (string, bool) {
	realPath := filepath.Join(appDataDir, name)
	if err := fsutil.EnsureDir(realPath); err != nil {
		return "", false
	}
	r, err := fsutil.AbsPath(realPath)
	if err != nil {
		return "", false
	}
	return r, true
}

func IsSubDirFileExist(name string) bool {
	r := GetDataDir()
	if r == "" {
		return false
	}
	return fsutil.FileExists(filepath.Join(r, name))
}

func GetSubDirFilePath(name string) string {
	r := GetDataDir()
	if r == "" {
		return ""
	}
	return filepath.Join(r, name)
}

func GetDataDir() string {
	wd, err := fsutil.WorkingDir()
	if err != nil {
		return ""
	}
	return filepath.Join(wd, appDataDir)
}
