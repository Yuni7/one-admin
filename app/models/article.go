package models

import (
	"gorm.io/gorm"
	"test-gin-admin/core"
	"test-gin-admin/schema"
	"test-gin-admin/util/structure"
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

type Articles []*Article

func (a Article) ToSchemaArticle() *schema.Article {
	item := new(schema.Article)
	structure.Copy(a, item)
	return item
}

func (a Articles) ToSchemaArticles() []*schema.Article {
	list := make([]*schema.Article, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaArticle()
	}
	return list
}
