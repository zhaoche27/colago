package dto

type PageQuery struct {
	*Query
	PageNum        int         `json:"page_num"`
	PageSize       int         `json:"page_size"`
	NeedTotalCount bool        `json:"need_total_count"`
	OrderDescs     []OrderDesc `json:"order_descs"`
}

func (pq PageQuery) Offset() int {
	if pq.PageNum > 0 {
		return (pq.PageNum - 1) * pq.PageSize
	}
	return 0
}
