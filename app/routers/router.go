package routers

import (
	"github.com/gin-gonic/gin"
	"test-gin-admin/handler"
)

func InitRouter() {

	r := gin.Default()
	//文章路由组 暂时没有加入 接口鉴权中间件
	articleGroup := r.Group("/article")
	{
		articleGroup.POST("/create-article", handler.CreateArticle)
		articleGroup.GET("/get-article-list", handler.GetArticleList)
		articleGroup.POST("/update-article", handler.UpdateArticle)
		articleGroup.DELETE(":id", handler.DeleteArticleById)
	}
	r.Run(":5560")
}
