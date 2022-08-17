package models

import (
	"github.com/volatiletech/null/v8"

	"time"
)

type Product struct {
	ID        string // uuid
	Name      string
	Price     float32
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
