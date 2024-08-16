package dto

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type PaginationResponse[T any] struct {
	Content    []*T `json:"content"`
	TotalItems int  `json:"totalItems"`
}
