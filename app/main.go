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
}
