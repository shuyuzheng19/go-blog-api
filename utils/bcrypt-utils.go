package utils

import (
	"gin-demo/myerr"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(str string) string {
	var encryptPassword, err = bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	myerr.MessageError(err, "密码加密失败")
	return string(encryptPassword)
}

func VerifyPassword(password string, encryptPassword string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
