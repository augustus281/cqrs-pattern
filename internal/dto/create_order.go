package dto

import "github.com/augustus281/cqrs-pattern/internal/order/models"

type CreateOderRequest struct {
	ShopItems       []*models.ShopItem `json:"shop_items,omitempty" validate:"required"`
	AccountEmail    string             `json:"account_email" validate:"required,email"`
	DeliveryAddress string             `json:"delivery_address" validate:"required"`
}
