package request

type CreateProduct struct {
	Name     string  `json:"name" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Quantity int64   `json:"quantity" validate:"required"`
}
