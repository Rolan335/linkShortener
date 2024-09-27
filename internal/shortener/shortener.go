package shortener

import (
	"crypto/sha1"
	"encoding/base64"
)

func Short(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	hashed := hash.Sum(nil)

	shortened := base64.RawURLEncoding.EncodeToString(hashed)
	return shortened[:6]
}
