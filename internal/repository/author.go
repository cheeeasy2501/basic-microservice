package repository

import (
	"basic-microservice/internal/domain/entity"
	"basic-microservice/pkg/database"
	"context"
	"time"
)

type IAuthorRepository interface {
	CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error)
	AssignBookByIds(ctx context.Context, ids []int) ([]entity.Author, error)
}

type AuthorRepository struct {
	db *database.Database
}

func newAuthorRepository(db *database.Database) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, Author entity.Author) (entity.Author, error) {
	Author.Id, Author.CreatedAt, Author.UpdatedAt = 1, time.Now(), time.Now()
	return Author, nil
}

func (r *AuthorRepository) AssignBookByIds(ctx context.Context, ids []int) ([]entity.Author, error) {
	authors := []entity.Author{}
	if len(ids) == 0 {
		return authors, nil
	}

	tx := r.db.GetSession(ctx)

	// todo assign id's
	//for _, v := range ids {
	//
	//	_ = v
	//	// todo: create dependencies between book and authors
	//}
	//todo return authors
	result := tx.Find(&authors, ids)

	if err := result.Error; err != nil {
		return authors, err
	}
	return authors, nil
}
