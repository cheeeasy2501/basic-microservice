package repository

import "gorm.io/gorm"

type Repositories struct {
	BookRepo IBookRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		BookRepo: newBookRepository(db),
	}
}
