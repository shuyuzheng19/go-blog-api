package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/response"
)

func CreateAccessToken(username string, nickName string) string {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(common.TokenExpire).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        username,
		Issuer:    nickName,
		Subject:   username,
	})

	accessToken, err := claims.SignedString([]byte(common.TokenEncrypt))

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "token生成失败!"))
	}

	return accessToken

}

func ParseTokenFormToUsername(token string) string {

	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.TokenEncrypt), nil
	})

	if err != nil || !claims.Valid {
		panic(response.NewGlobalException(response.AUTHENTICATION, "用户认证失败,请重新登录!"))
	}

	return claims.Claims.(*jwt.StandardClaims).Subject
}
