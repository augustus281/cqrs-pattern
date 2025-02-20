package mongomodels

import (
	"fmt"

	orderservice "github.com/augustus281/cqrs-pattern/api"
)

type ShopItem struct {
	ID          string  `json:"id" bson:"id,omitempty"`
	Title       string  `json:"title" bson:"title,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	Quantity    uint64  `json:"quantity" bson:"quantity,omitempty"`
	Price       float64 `json:"price" bson:"price,omitempty"`
}

func (s *ShopItem) String() string {
	return fmt.Sprintf("ID: {%s}, Title: {%s}, Description: {%s}, Quantity: {%v}, Price: {%v},",
		s.ID,
		s.Title,
		s.Description,
		s.Quantity,
		s.Price,
	)
}

func (s *ShopItem) ToProto() *orderservice.ShopItem {
	return &orderservice.ShopItem{
		Id:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		Quantity:    s.Quantity,
		Price:       s.Price,
	}
}

func ShopItemToProto(shopItem *ShopItem) *orderservice.ShopItem {
	return &orderservice.ShopItem{
		Id:          shopItem.ID,
		Title:       shopItem.Title,
		Description: shopItem.Description,
		Quantity:    shopItem.Quantity,
		Price:       shopItem.Price,
	}
}

func ShopItemFromProto(shopItem *orderservice.ShopItem) *ShopItem {
	return &ShopItem{
		ID:          shopItem.Id,
		Title:       shopItem.Title,
		Description: shopItem.Description,
		Quantity:    shopItem.Quantity,
		Price:       shopItem.Price,
	}
}

func ShopItemsToProto(shopItems []*ShopItem) []*orderservice.ShopItem {
	items := make([]*orderservice.ShopItem, 0, len(shopItems))
	for _, item := range shopItems {
		items = append(items, ShopItemToProto(item))
	}
	return items
}

func ShopItemsFromProto(shopItems []*orderservice.ShopItem) []*ShopItem {
	items := make([]*ShopItem, 0, len(shopItems))
	for _, item := range shopItems {
		items = append(items, ShopItemFromProto(item))
	}
	return items
}
