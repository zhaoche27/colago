package dto

type OrderDesc struct {
	Col string `json:"col"`
	Asc bool   `json:"asc"`
}

func NewOrderDescWithoutAsc(col string) OrderDesc {
	return NewOrderDesc(col, true)
}

func NewOrderDesc(col string, asc bool) OrderDesc {
	return OrderDesc{Col: col, Asc: asc}
}
