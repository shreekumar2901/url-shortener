package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`       // Stored as a hash
	Role      string    `gorm:"default:user"`   // e.g., "admin" or "user"
	CreatedAt time.Time `gorm:"autoCreateTime"` // Timestamp when the user was created
}
