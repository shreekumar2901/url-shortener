package domain

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	ID     string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Url    string    `gorm:"not null;uniqueIndex:idx_user_url_short"`
	Short  string    `gorm:"not null;uniqueIndex:idx_user_url_short"`
	Expiry time.Time `gorm:"not null"`

	UserID string `gorm:"type:uuid;not null;uniqueIndex:idx_user_url_short"`                // Foreign key column
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key relationship
}
