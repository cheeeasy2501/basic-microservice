package repository

import (
	"basic-microservice/pkg/database"
)

type Repositories struct {
	BookRepo   IBookRepository
	AuthorRepo IAuthorRepository
}

func NewRepositories(db *database.Database) *Repositories {
	return &Repositories{
		BookRepo:   newBookRepository(db),
		AuthorRepo: newAuthorRepository(db),
	}
}
