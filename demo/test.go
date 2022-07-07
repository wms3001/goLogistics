package demo

import "github.com/gin-gonic/gin"

func Message(c *gin.Context) {
	c.JSON(200, "I not need auth")
}

func Message1(c *gin.Context) {
	c.JSON(200, "I need auth")
}
