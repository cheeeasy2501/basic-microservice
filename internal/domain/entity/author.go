package entity

import (
	"gorm.io/gorm"
	"time"
)

type AuthorEntity struct {
	Id        uint64         `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	deletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *AuthorEntity) TableName() string {
	return "authors"
}
