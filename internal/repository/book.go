package repository

import (
	"basic-microservice/internal/entity"
	"basic-microservice/pkg/database"
	"context"
	"time"
)

type IBookRepository interface {
	CreateBook(ctx context.Context, book entity.BookEntity) (entity.BookEntity, error)
}

type BookRepository struct {
	db *database.Database
}

func newBookRepository(db *database.Database) *BookRepository {
	return &BookRepository{db: db}
}

// return entity.Book
func (r *BookRepository) CreateBook(ctx context.Context, book entity.BookEntity) (entity.BookEntity, error) {
	// todo: create book and return id, created_at and updated_at fields
	book.Id, book.CreatedAt, book.UpdatedAt = 1, time.Now(), time.Now()
	return book, nil
}
