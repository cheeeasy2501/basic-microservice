package service

import (
	"basic-microservice/internal/entity"
	repos "basic-microservice/internal/repository"
)

type IBookService interface {
	GetBooks()
	GetBook()
	CreateBook(ent entity.BookEntity) (entity.BookEntity, error)
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

func (s *BookService) CreateBook(book entity.BookEntity) (entity.BookEntity, error) {
	s.bookRepo.CreateBook() // todo: create book and return id's
	book.Id = 1             // mock
	return book, nil
}

func (s *BookService) UpdateBook() {

}

func (s *BookService) DeleteBook() {

}
