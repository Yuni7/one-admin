package service

import (
	"test-gin-admin/core"
	"test-gin-admin/models"
	"test-gin-admin/schema"
	"test-gin-admin/util"
)

func Query(params schema.ArticleQueryParam, opts ...schema.ArticleQueryOptions) (*schema.ArticleQueryResult, error) {
	var opt schema.ArticleQueryOptions
	db := core.GetDB().Model(new(schema.Article))

	if len(opts) > 0 {
		opt = opts[0]
	}
	if v := params.Title; v != "" {
		db = db.Where("title=?", v)
	}
	if v := params.Content; v != "" {
		db = db.Where("content=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}
	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}
	var list models.Articles
	pr, err := util.WrapPageQuery(db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
	}
	qr := &schema.ArticleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaArticles(),
	}
	return qr, nil
}

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
