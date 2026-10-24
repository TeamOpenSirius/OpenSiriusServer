package repositories

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

func NewToken(prefix string) (string, error) {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		prefix = "tok"
	}
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s", prefix, base64.RawURLEncoding.EncodeToString(buf)), nil
}
