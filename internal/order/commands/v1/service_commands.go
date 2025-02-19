package v1

type OrderCommands struct {
	CreateOrder           CreateOrder
	OrderPaid             PayOrder
	SubmitOrder           SubmitOrder
	UpdateOrder           UpdateShoppingCart
	CancelOrder           CancelOrder
	CompleteOrder         CompleteOrder
	ChangeDeliveryAddress ChangeDeliveryAddress
}

func NewOrderCommands(
	createOrder CreateOrder,
	orderPaid PayOrder,
	submitOrder SubmitOrder,
	updateOrder UpdateShoppingCart,
	cancelOrder CancelOrder,
	completeOrder CompleteOrder,
	changeDechangeDeliveryAddress ChangeDeliveryAddress,
) *OrderCommands {
	return &OrderCommands{
		CreateOrder:           createOrder,
		OrderPaid:             orderPaid,
		SubmitOrder:           submitOrder,
		UpdateOrder:           updateOrder,
		CancelOrder:           cancelOrder,
		CompleteOrder:         completeOrder,
		ChangeDeliveryAddress: changeDechangeDeliveryAddress,
	}
}
