package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func IsEmpty(str string) bool {
	return str == "" || strings.TrimSpace(str) == ""
}

func AnyEmpty(strings ...string) bool {
	for _, str := range strings {
		if IsEmpty(str) {
			return true
		}
	}
	return false
}

func ToMd5String(bytes []byte) string {
	hash := md5.New()

	hash.Write(bytes)

	md5String := hex.EncodeToString(hash.Sum(nil))

	return md5String
}
