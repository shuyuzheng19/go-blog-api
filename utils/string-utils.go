package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"mime/multipart"
	"strconv"
)

func CreateRandomNumber() string {
	return strconv.Itoa(rand.Intn(900000) + 100000)
}

func GetFileMd5(file multipart.File) string {
	hash := md5.New()

	io.Copy(hash, file)

	md5 := hex.EncodeToString(hash.Sum(nil))

	file.Close()

	return md5
}
