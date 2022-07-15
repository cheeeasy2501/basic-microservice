package repository

import (
	"basic-microservice/internal/domain/entity"
	"basic-microservice/pkg/database"
	"context"
	"time"
)

type IAuthorRepository interface {
	CreateAuthor(ctx context.Context, author entity.AuthorEntity) (entity.AuthorEntity, error)
	AssignBookByIds(ctx context.Context, ids []int) ([]entity.AuthorEntity, error)
}

type AuthorRepository struct {
	db *database.Database
}

func newAuthorRepository(db *database.Database) *AuthorRepository {
	return &AuthorRepository{db: db}
}

// return entity.Author
func (r *AuthorRepository) CreateAuthor(ctx context.Context, Author entity.AuthorEntity) (entity.AuthorEntity, error) {
	// todo: create Author and return id, created_at and updated_at fields
	Author.Id, Author.CreatedAt, Author.UpdatedAt = 1, time.Now(), time.Now()
	return Author, nil
}

func (r *AuthorRepository) AssignBookByIds(ctx context.Context, ids []int) ([]entity.AuthorEntity, error) {
	tx := r.db.GetSession(ctx)

	authors := []entity.AuthorEntity{}
	if len(ids) == 0 {
		return authors, nil
	}

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
