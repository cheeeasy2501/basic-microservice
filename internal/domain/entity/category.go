package entity

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        uint64         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	deletedAt gorm.DeletedAt `gorm:"index"`
}
