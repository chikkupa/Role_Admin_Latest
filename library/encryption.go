package library

import(
	"crypto/sha1"
    "encoding/hex"
)

func ConvertToSha1(value string) string{
    h := sha1.New()
    h.Write([]byte(value))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

    return sha1_hash
}