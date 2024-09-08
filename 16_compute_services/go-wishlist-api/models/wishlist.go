package models

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title      string         `json:"title"`
	IsAchieved bool           `json:"is_achieved"`
}
