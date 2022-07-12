package repository

import "gorm.io/gorm"

type IBookRepository interface {
	CreateBook()
}

type BookRepository struct {
	db *gorm.DB
}

func newBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// return entity.Book
func (r *BookRepository) CreateBook() {

}
