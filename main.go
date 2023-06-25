package main

import (
	"fmt"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case common.Result:
				c.JSON(http.StatusOK, common.IError(t.Code, t.Message))
				c.Abort()
				config.LOGGER.Error("自定义错误=>", zap.Int("code", t.Code), zap.String("message", t.Message))
				break
			default:
				c.JSON(http.StatusOK, common.Error("服务器异常喽........"))
				c.Abort()
				config.LOGGER.Error("服务器异常消息=>")
				break
			}
		}
	}()
	c.Next()
}

func Corn(context *gin.Context) {
	defer func() {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

	}()
}

func SetupConfig() {
	config.LOGGER = config.LoadLogger()

	defer config.LOGGER.Sync()

	config.LoadConfig()

	config.LoadDbConfig()

	config.LoadRedisConfig()

	config.LoadMeiliSearchConfig()
}

func main() {

	SetupConfig()

	var server = gin.Default()

	server.Use(Corn, Recover)

	var router = routers.NewRouters(server.RouterGroup)

	router.SetupRouter()

	var addr = fmt.Sprintf(":%d", config.CONFIG.Server.Port)

	server.Run(addr)

	router.EnableCronJob()

}
