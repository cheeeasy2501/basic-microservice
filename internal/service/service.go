package service

import repo "basic-microservice/internal/repository"

type Services struct {
	BookService IBookService
}

func NewServices(repos *repo.Repositories) *Services {
	return &Services{
		BookService: NewBookService(repos.BookRepo),
	}
}
