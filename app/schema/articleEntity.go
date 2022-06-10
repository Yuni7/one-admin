package schema

import "time"

type Article struct {
	ID        uint64    `json:"id,string"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
