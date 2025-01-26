package entities

import (
	"time"
)

type Product struct {
	Id          int
	Name        string
	Category Category
	Stock int
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}