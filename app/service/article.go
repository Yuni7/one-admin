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
