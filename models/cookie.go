package models

// Cookie storage backed by AES-256-GCM with Base64 encoding.
// Domain, secure flag, max age, and the encryption key are read from conf/app.ini.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

type ginCookie struct{}

type cookieConfig struct {
	key      []byte
	domain   string
	secure   bool
	httpOnly bool
	maxAge   int
}

func loadCookieConfig() cookieConfig {
	cfg := cookieConfig{
		key:      []byte("0123456789abcdef0123456789abcdef"),
		domain:   "127.0.0.1",
		secure:   false,
		httpOnly: true,
		maxAge:   3600,
	}

	file, err := ini.Load("./conf/app.ini")
	if err != nil {
		return cfg
	}

	section := file.Section("cookie")
	key := []byte(section.Key("key").String())
	if len(key) > 0 {
		if len(key) != 32 {
			sum := sha256.Sum256(key)
			key = sum[:]
		}
		cfg.key = key
	}
	if domain := section.Key("domain").String(); domain != "" {
		cfg.domain = domain
	}
	cfg.secure = section.Key("secure").MustBool(false)
	cfg.httpOnly = section.Key("httpOnly").MustBool(true)
	cfg.maxAge = section.Key("maxAge").MustInt(3600)

	return cfg
}

func encryptCookie(plain []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	sealed := gcm.Seal(nonce, nonce, plain, nil)
	return base64.URLEncoding.EncodeToString(sealed), nil
}

func decryptCookie(value string, key []byte) ([]byte, error) {
	data, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, errors.New("cookie: ciphertext too short")
	}

	nonce := data[:gcm.NonceSize()]
	ciphertext := data[gcm.NonceSize():]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cfg := loadCookieConfig()
	encData, err := encryptCookie(bytes, cfg.key)
	if err != nil {
		return err
	}

	domain := cfg.domain
	if domain == "" {
		domain = c.Request.Host
	}
	c.SetCookie(key, encData, cfg.maxAge, "/", domain, cfg.secure, cfg.httpOnly)
	return nil
}

func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err := c.Cookie(key)
	if err != nil || valueStr == "" {
		return false
	}

	cfg := loadCookieConfig()
	decData, err := decryptCookie(valueStr, cfg.key)
	if err != nil {
		return false
	}

	return json.Unmarshal(decData, obj) == nil
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	cfg := loadCookieConfig()
	domain := cfg.domain
	if domain == "" {
		domain = c.Request.Host
	}
	c.SetCookie(key, "", -1, "/", domain, cfg.secure, cfg.httpOnly)
	return true
}

var Cookie = &ginCookie{}
