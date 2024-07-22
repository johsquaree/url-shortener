package shortener

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateShortURL(originalURL string) string {
	h := sha1.New()
	h.Write([]byte(originalURL))
	bs := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bs)[:8] // İlk 8 karakteri al
}
