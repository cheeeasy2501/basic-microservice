package repository

import (
	"basic-microservice/internal/domain/entity"
	"basic-microservice/pkg/database"
	"context"
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
	db := r.db.GetSession(ctx)
	result := db.WithContext(ctx).Create(&book)
	if err := result.Error; err != nil {
		return entity.BookEntity{}, err
	}

	return book, nil
}
