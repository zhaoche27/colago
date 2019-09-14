package dto

type PageQuery struct {
	*Query
	PageNum        int
	PageSize       int
	NeedTotalCount bool
	OrderDescs     []OrderDesc
}

func (pq PageQuery) Offset() int {
	if pq.PageNum > 0 {
		return (pq.PageNum - 1) * pq.PageSize
	}
	return 0
}
