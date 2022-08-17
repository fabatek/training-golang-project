package request

type OrderItem struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int64  `json:"quantity" validate:"required"`
}
