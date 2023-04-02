package models

// task1

type SendProduct struct {
	SenderId   int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type OrderTotalSum struct {
	OrderId       int    `json:"order_id"`
	PromocodeName string `json:"promocode_name"`
}
