package service

import (
	"test-gin-admin/core"
	"test-gin-admin/schema"
)

func CreatePostService(item schema.Article) (ok bool) {
	db := core.GetDB()
	tx := db.Create(&item)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
func GetArticleListService() (articles []*schema.Article) {
	db := core.GetDB()
	db.Find(&articles)
	return articles
}
func Update(item schema.Article) (ok bool) {
	db := core.GetDB()
	tx := db.Model(&item).Updates(item)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
func DeleteArticleById(id uint64) bool {
	db := core.GetDB()
	tx := db.Delete(id)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
