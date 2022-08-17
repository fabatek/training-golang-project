package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type User struct {
	ID        string // uuid
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
