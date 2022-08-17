package response

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	OrderItem []OrderItem `json:"order_item"`
	CreateAt  string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}
