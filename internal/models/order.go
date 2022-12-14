package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Order struct {
	ID        string // uuid
	UserID    string
	SumPrice  float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
