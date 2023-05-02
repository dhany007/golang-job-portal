package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

func SHA1(val string) string {
	sha1 := sha1.New()

	_, err := sha1.Write([]byte(val))
	if err != nil {
		return ""
	}

	return hex.EncodeToString(sha1.Sum(nil))
}

func StringToInt(val string, def int64) int64 {
	r, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		r = int64(def)
	}

	return r
}
