package domain

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	ID     string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Url    string    `gorm:"unique;not null"`
	Short  string    `gorm:"unique;not null"`
	Expiry time.Time `gorm:"not null"`
}
