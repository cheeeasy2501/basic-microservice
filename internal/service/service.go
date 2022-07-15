package service

import (
	repo "basic-microservice/internal/repository"
	"basic-microservice/pkg/database"
)

type Services struct {
	BookService IBookService
}

func NewServices(db *database.Database, repos *repo.Repositories) *Services {
	return &Services{
		BookService: NewBookService(db, repos.BookRepo, repos.AuthorRepo),
	}
}
