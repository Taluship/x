package pclient

type PaginationResponse[T any] struct {
	Items      []*T `json:"items"`
	Page       int  `json:"page"`
	PerPage    int  `json:"perPage"`
	TotalItems int  `json:"totalItems"`
	TotalPages int  `json:"totalPages"`
}
