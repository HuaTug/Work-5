package utils

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
)

func MD5(str string) string {
	h := md5.New() //nolint:gosec
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
