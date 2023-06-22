package middleware

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authorized(role string) gin.HandlerFunc {
	return func(context *gin.Context) {

		var tokenHeader = context.GetHeader(common.TOKEN_HEADER)

		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, common.TOKEN_TYPE) {
			context.JSON(http.StatusOK, common.NO_LOGIN)
			context.Abort()
			return
		}

		var token = strings.Replace(tokenHeader, common.TOKEN_TYPE, "", 1)

		var username = utils.ParseTokenToUsername(token)

		if username == "" {
			context.JSON(http.StatusOK, common.USER_NOT_FOUNT)
			context.Abort()
			return
		}

		var user = FinByUsername(username)

		if user.Id == 0 {
			context.JSON(http.StatusOK, common.USER_NOT_FOUNT)
			context.Abort()
			return
		}

		var redisToken = utils.GetCurrentUserToken(username)

		if redisToken != token {
			context.JSON(http.StatusOK, common.AUTHENTICATE_ERROR)
			context.Abort()
			return
		}

		var isAuth = false

		var roleName = user.Role.Name

		if role == common.USER_ROLE {
			isAuth = true
		} else if role == common.SUPER_ROLE && roleName == common.SUPER_ROLE {
			isAuth = true
		} else if role == common.USER_ROLE && roleName == common.USER_ROLE {
			isAuth = true
		}

		if isAuth {
			context.Set("user", user)
			context.Next()
		} else {
			context.JSON(http.StatusOK, common.AUTHORIZED_ERROR)
			context.Abort()
		}

	}
}

func FinByUsername(username string) (user models.User) {
	var result = config.REDIS.Get(common.USER_INFO_KEY + username).Val()

	if result == "" {
		config.DB.Model(&models.User{}).Preload("Role").First(&user, "username = ?", username)
		if user.Id > 0 {
			var userJson, _ = json.Marshal(&user)
			config.REDIS.Set(common.USER_INFO_KEY+user.Username, userJson, common.USER_INFO_EXPIRE)
			return user
		}
	} else {
		json.Unmarshal([]byte(result), &user)
	}

	return user
}
