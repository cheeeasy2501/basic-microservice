package service

import (
	"basic-microservice/internal/entity"
	repos "basic-microservice/internal/repository"
	"context"
)

type IBookService interface {
	GetBooks()
	GetBook()
	CreateBook(ctx context.Context, ent entity.BookEntity) (entity.BookEntity, error)
	UpdateBook()
	DeleteBook()
}

type BookService struct {
	bookRepo repos.IBookRepository
}

func NewBookService(bookRepo repos.IBookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

// return entity.Book
func (s *BookService) GetBooks() {

}

func (s *BookService) GetBook() {

}

func (s *BookService) CreateBook(ctx context.Context, book entity.BookEntity) (entity.BookEntity, error) {
	book, err := s.bookRepo.CreateBook(ctx, book) // todo: create book and return id's
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *BookService) UpdateBook() {

}

func (s *BookService) DeleteBook() {

}
