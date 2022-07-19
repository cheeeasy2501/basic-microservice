package repository

import (
	"basic-microservice/internal/domain/entity"
	"basic-microservice/pkg/database"
	"context"
)

type IBookRepository interface {
	CreateBook(ctx context.Context, book entity.Book) (entity.Book, error)
}

type BookRepository struct {
	db *database.Database
}

func newBookRepository(db *database.Database) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	db := r.db.GetSession(ctx)
	result := db.WithContext(ctx).Create(&book)
	if err := result.Error; err != nil {
		return entity.Book{}, err
	}

	return book, nil
}
