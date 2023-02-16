package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/repository"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type filterChain struct {
	Path string
	Role string
}

var PREFIX = "/api/v1"

var filterChains = []filterChain{
	{
		Path: PREFIX + "/user/get",
		Role: common.RoleUser,
	},
	{
		Path: PREFIX + "/upload/file",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/file/current_user/list",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/file/current_user/delete",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/file/delete",
		Role: common.RoleSuperAdmin,
	},
	{
		Path: PREFIX + "/blog/user_save",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/blog/is_like",
		Role: common.User,
	},
	{
		Path: PREFIX + "/blog/like",
		Role: common.User,
	},
	{
		Path: PREFIX + "/blog/get/user_save",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/tag/add",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/category/add",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/topic/add",
		Role: common.RoleAdmin,
	},
	{
		Path: PREFIX + "/user/logout",
		Role: common.RoleUser,
	},
	{
		Path: PREFIX + "/blog/add",
		Role: common.RoleAdmin,
	},
}

func findWhiteUris(path string) int {

	for i, chain := range filterChains {
		if path == chain.Path || strings.Contains(path, chain.Path) {
			return i
		}
	}

	return -1
}

func GlobalFilter() gin.HandlerFunc {

	return func(context *gin.Context) {

		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)

			context.Next()

			return
		}

		path := context.Request.URL.Path

		index := findWhiteUris(path)

		if index == -1 {
			context.Next()
			return
		}

		var tokenHeader = context.GetHeader(common.TokenHeader)

		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, common.TokenPrefix) {
			context.JSON(200, response.FAILURE(response.AUTHENTICATION, "你还未登录,请登录后重试"))
			context.Abort()
			return
		}

		var token = tokenHeader[len(common.TokenPrefix):]

		username := utils.ParseTokenFormToUsername(token)

		resultToken := config.Redis.Get(common.UserToken + ":" + utils.ToMd5String([]byte(username))).Val()

		if resultToken != token {
			context.JSON(200, response.FAILURE(response.AUTHENTICATION, "身份验证失败,请重新登录"))
			context.Abort()
			return
		}

		user := repository.FindByUsername(username)

		if user.Id == 0 {
			context.JSON(200, response.FAILURE(response.AUTHENTICATION, "身份验证失败,请重新登录"))
		}

		chain := filterChains[index]

		if hasRole(chain.Role, user.Role.Name) {
			context.Set("user", user)
			context.Next()
		} else {
			context.JSON(200, response.FAILURE(response.AUTHORIZATION, "你的权限不够,请联系管理员"))
			context.Abort()
		}

	}
}

func hasRole(chainRole string, currentRole string) bool {

	if chainRole == currentRole || chainRole == common.RoleUser {
		return true
	}

	if chainRole == common.RoleAdmin && (currentRole == common.RoleAdmin || currentRole == common.RoleSuperAdmin) {
		return true
	}

	return false
}
