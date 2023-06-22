package utils

import (
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/myerr"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateAccessToken(username string, nickName string) string {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(common.TOKEN_EXPIRE).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        username,
		Issuer:    nickName,
		Subject:   username,
	})

	accessToken, err := claims.SignedString([]byte(common.TOKEN_ENCRYPT))

	if err != nil {
		myerr.PanicError(common.TOKEN_GENERATE_ERROR)
	}

	return accessToken

}

func ParseTokenToUsername(token string) string {

	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.TOKEN_ENCRYPT), nil
	})

	if err != nil || !claims.Valid {
		return ""
	}

	return claims.Claims.(*jwt.StandardClaims).Subject
}

func GetCurrentUserToken(username string) string {
	return config.REDIS.Get(common.TOKEN_KEY + username).Val()
}
