package dto

import "github.com/augustus281/cqrs-pattern/internal/order/models"

type UpdateShopItemsReqDTO struct {
	ShopItems []*models.ShopItem `json:"shop_items,omitempty" validate:"required"`
}
