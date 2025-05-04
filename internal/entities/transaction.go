package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string    `gorm:"type:varchar(50);not null"`
	Status        string    `gorm:"type:varchar(20);not null"` // verified, completed, failed
	Balance       float64   `gorm:"type:decimal(10,2)"`
	ExpiresAt     time.Time `gorm:"type:timestamp;not null"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp"`
}
