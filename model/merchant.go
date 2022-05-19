package model

import (
	"gorm.io/gorm"
	"time"
)

type Merchant struct {
	gorm.Model
	ID        string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
