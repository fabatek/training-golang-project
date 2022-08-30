package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type OrderItem struct {
	ID        string // uuid
	ProductID string
	Quantity  int64
	Price     float32
	OrderID   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
