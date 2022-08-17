package response

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Quantity  int64   `json:"quantity"`
	CreateAt  string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
