package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type SubscriptionWithoutId struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" json:"id"`
	Price     float32   `json:"price" validate:"required"`
	Frequency int       `json:"frequency"`
	Plan      string    `json:"plan" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Subscription struct {
	UserId uuid.UUID `gorm:"primaryKey;type:uuid;" json:"user" validate:"required"`
	SubscriptionWithoutId
}
