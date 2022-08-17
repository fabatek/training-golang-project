package request

type CreateOrder struct {
	UserID    string      `json:"user_id" validate:"required"`
	OrderItem []OrderItem `json:"order_item"`
}
