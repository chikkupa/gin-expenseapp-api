package library

import(
	"crypto/md5"
    "encoding/hex"
)

// GetMd5 Converts a string to md5 hash
func GetMd5(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}