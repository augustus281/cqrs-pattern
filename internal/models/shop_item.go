package models

type ShopItem struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Quantity    uint64  `json:"quantity,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
