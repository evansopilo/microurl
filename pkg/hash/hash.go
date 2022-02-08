package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(String string) string {

	hash := sha256.New()

	hash.Write([]byte(String))

	url_hash := hex.EncodeToString(hash.Sum(nil))[0:7]

	return url_hash
}
