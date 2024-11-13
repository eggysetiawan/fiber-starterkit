package dto

type WithPaginationResponse[T any] struct {
	Meta MetaResponse `json:"meta"`
	Data []T          `json:"data"`
}

type MetaResponse struct {
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	Limit       int    `json:"per_page"`
	Total       int    `json:"total_record"`
	LastQueryAt string `json:"last_query_at"`
	Offset      int    `json:"offset"`
}
