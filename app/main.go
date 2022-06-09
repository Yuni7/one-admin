package main

import (
	"github.com/gin-gonic/gin"
	"test-gin-admin/routers"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routers.InitRouter()
	//test
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "success",
	//	})
	//})
	//
	//r.Run(":8080")
}
