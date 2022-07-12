package repository

import "gorm.io/gorm"

type Repositories struct {
	bookRepo IBookRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		bookRepo: newBookRepository(db),
	}
}
