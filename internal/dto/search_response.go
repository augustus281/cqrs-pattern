package dto

type OrderSearchResponseDTO struct {
	Pagination Pagination         `json:"pagination"`
	Orders     []OrderResponseDto `json:"orders"`
}

type Pagination struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"has_more"`
}
