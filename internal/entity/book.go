package entity

import (
	"gorm.io/gorm"
	"time"
)

type BookEntity struct {
	Id          uint64         `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
