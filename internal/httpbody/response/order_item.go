package response

type OrderItem struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}
