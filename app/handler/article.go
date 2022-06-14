package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-gin-admin/schema"
	"test-gin-admin/service"
	"test-gin-admin/types"
	"test-gin-admin/util"
)

func Query(c *gin.Context) {
	var params schema.ArticleQueryParam
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params.Pagination = true
	result, err := service.Query(params, schema.ArticleQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("id", schema.OrderByDESC)),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.Success(c, result)

}

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
func UpdateArticle(c *gin.Context) {
	var article schema.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok := service.Update(article)
	if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
	util.Success(c, nil)
}
func DeleteArticleById(c *gin.Context) {
	key := "id"
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok := service.DeleteArticleById(id)
	if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
	util.Success(c, nil)
}
