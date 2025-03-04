package mappers

import (
	orderService "github.com/augustus281/cqrs-pattern/api"
	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
)

type ShopItem struct {
	ID          string  `json:"id" bson:"id,omitempty"`
	Title       string  `json:"title" bson:"title,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	Quantity    uint64  `json:"quantity" bson:"quantity,omitempty"`
	Price       float64 `json:"price" bson:"price,omitempty"`
}

func ShopItemToProto(shopItem *ShopItem) *orderService.ShopItem {
	return &orderService.ShopItem{
		Id:          shopItem.ID,
		Title:       shopItem.Title,
		Description: shopItem.Description,
		Quantity:    shopItem.Quantity,
		Price:       shopItem.Price,
	}
}

func ShopItemFromProto(shopItem *orderService.ShopItem) *ShopItem {
	return &ShopItem{
		ID:          shopItem.Id,
		Title:       shopItem.Title,
		Description: shopItem.Description,
		Quantity:    shopItem.Quantity,
		Price:       shopItem.Price,
	}
}

func ShopItemResponseFromModel(item *models.ShopItem) dto.ShopItem {
	return dto.ShopItem{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}

func ShopItemsResponseFromModels(items []*models.ShopItem) []dto.ShopItem {
	shopItems := make([]dto.ShopItem, 0, len(items))
	for _, item := range items {
		shopItems = append(shopItems, ShopItemResponseFromModel(item))
	}
	return shopItems
}

func ShopItemsToProto(shopItems []*ShopItem) []*orderService.ShopItem {
	items := make([]*orderService.ShopItem, 0, len(shopItems))
	for _, item := range shopItems {
		items = append(items, ShopItemToProto(item))
	}
	return items
}

func ShopItemsFromProto(shopItems []*orderService.ShopItem) []*ShopItem {
	items := make([]*ShopItem, 0, len(shopItems))
	for _, item := range shopItems {
		items = append(items, ShopItemFromProto(item))
	}
	return items
}
