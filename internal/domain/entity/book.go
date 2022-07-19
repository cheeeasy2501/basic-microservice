package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id              uint64         `json:"id" gorm:"primaryKey"`
	Isbn            string         `json:"isbn"`
	Status          string         `json:"status"`
	Title           string         `json:"title"`
	LanguageLevelId int            `json:"languageLevelId"`
	Description     string         `json:"description"`
	Link            string         `json:"link"`
	CoverPath       string         `json:"coverPath"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	deletedAt       gorm.DeletedAt `gorm:"index"`
}

func (b *Book) TableName() string {
	return "books"
}
