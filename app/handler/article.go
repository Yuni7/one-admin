package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin-admin/schema"
	"test-gin-admin/service"
	"test-gin-admin/types"
	"test-gin-admin/util"
)

func CreateArticle(c *gin.Context) {
	var article schema.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok := service.CreatePostService(article)
	if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	}
	util.Success(c, nil)
}
func GetArticleList(c *gin.Context) {
	article := service.GetArticleListService()

	util.Success(c, article)
}
