package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"net/http"
	"vs-blog-api/config"
	"vs-blog-api/filter"
	"vs-blog-api/manager"
	"vs-blog-api/response"
	"vs-blog-api/router"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case response.GlobalException:
				c.JSON(200, response.FAILURE(t.Code, t.Message))
				c.Abort()
				break
			default:
				c.JSON(200, response.FAILURE(http.StatusInternalServerError, "服务器异常"))
				c.Abort()
				break
			}
		}
	}()
	c.Next()
}

func InitConfig() {
	config.InitConfig()

	config.InitDb()

	config.InitRedis()

	config.InitElastic()
}

func init() {

	InitConfig()

	var system = manager.NewSystemManager()

	system.DbToElastic()

	c := cron.New()

	c.AddFunc("0 0 12 * * ?", func() {
		system.RedisEyeAndLikeCountToDb()
	})

	c.AddFunc("0 0 1 * * ?", func() {
		system.RedisEyeAndLikeCountToDb()
	})
	//c.Start()
}

func main() {

	var server = gin.Default()

	server.Use(Recover, filter.GlobalFilter())

	var userRouter = router.NewUserRouter(server.RouterGroup)

	userRouter.InitRouter()

	server.Run(":8888")

}
