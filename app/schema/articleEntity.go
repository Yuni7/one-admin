package schema

import (
	"time"
)

type Article struct {
	ID        uint64    `json:"id,string"`
	Title     string    `json:"title" `
	Content   string    `json:"content"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Articles 文章列表
type Articles []*Article

// ArticleQueryParam 查询条件
type ArticleQueryParam struct {
	PaginationParam
	Title      string `from:"title"`
	Content    string `from:"content"`
	QueryValue string `from:"queryValue"`
}

// ArticleQueryOptions 查询可选参数项
type ArticleQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []*string
}

// ArticleQueryResult 查询结果
type ArticleQueryResult struct {
	Data       Articles
	PageResult *PaginationResult
}
