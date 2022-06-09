package models

import (
	"gorm.io/gorm"
	"test-gin-admin/core"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	State   int    `gorm:"column:state"`
}

func (a *Article) Article() string {
	return "article"
}
func init() {
	core.Db.AutoMigrate(&Article{})
}
