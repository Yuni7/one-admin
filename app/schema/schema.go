package schema

type StatusText string

func (t StatusText) String() string {
	return string(t)
}

const (
	OKStatus    StatusText = "OK"
	ErrorStatus StatusText = "ERROR"
	FailStatus  StatusText = "FAIL"

	OrderByASC OrderDirection = iota + 1
	OrderByDESC
)

type OrderDirection int

type PaginationResult struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
}

type PaginationParam struct {
	Pagination bool `form:"-"`
	OnlyCount  bool `form:"-"`
	Current    int  `form:"current,default=1"`
	PageSize   int  `form:"pageSize,default=10" binding:"max=100"`
}

func (a PaginationParam) GetCurrent() int {
	return a.Current
}
func (a PaginationParam) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize <= 0 {
		pageSize = 100
	}
	return pageSize
}

type OrderField struct {
	Key       string
	Direction OrderDirection
}

func NewOrderFields(orderField ...*OrderField) []*OrderField {
	return orderField
}
func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}
