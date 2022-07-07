package route

import (
	"github.com/gin-gonic/gin"
	"logistics/demo"
	"logistics/oauth2"
	"logistics/service"
)

func Route(g *gin.Engine) {
	var c *gin.Context
	auth := g.Group("/auth")
	{
		auth.GET("/token", oauth2.TokenRequest)
		auth.GET("/credentials", oauth2.Credentials)
	}
	de := g.Group("/demo")
	{
		de.GET("/message", demo.Message)
	}
	de1 := g.Group("/demo1")
	{
		//权限认证中间件
		de1.Use(oauth2.AuthValidate(c))
		de1.GET("/message", demo.Message1)
	}
	order := g.Group("/order")
	{
		//权限认证中间件
		//order.Use(oauth2.AuthValidate(c))
		order.POST("/addOrder", service.AddOrder)
		order.POST("/addOrders", service.AddOrders)
		order.GET("/getOrderByOrderId", service.GetOrderByOrderId)
		order.GET("/getOrderById", service.GetOrderById)
		order.POST("/getOrder", service.GetOrder)
		order.POST("/getOrders", service.GetOrders)
		order.POST("/updateOrder", service.UpdateOrder)
		order.POST("/updateOrders", service.UpdateOrders)
		order.POST("/deleteOrder", service.DeleteOrder)
		order.POST("/deleteOrders", service.DeleteOrders)
		order.POST("/countOrder", service.CountOrder)
		order.POST("/distinctOrder", service.DistinctOrder)
	}

	amqp := g.Group("/amqp")
	{
		amqp.POST("/inTestQueue", service.InTestQueue)
		amqp.GET("/outTestQueue", service.OutTestQueue)
	}
}
