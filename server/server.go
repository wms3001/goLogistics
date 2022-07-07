package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"logistics/cron"
	"logistics/route"
	"os"
)

func Start() {
	//读取服务配置
	work, _ := os.Getwd()
	viper.SetConfigName("server")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/conf")
	viper.ReadInConfig()
	cron.TaskStart()
	g := gin.Default()
	//初始化路由
	route.Route(g)
	g.Run(":" + viper.GetString("server.port"))
}
