package service

import (
	"basic-microservice/internal/domain/aggregate"
	repos "basic-microservice/internal/repository"
	"basic-microservice/pkg/database"
	"context"
)

type IBookService interface {
	GetBooks()
	GetBook()
	CreateBook(ctx context.Context, ent aggregate.CreateBook) (aggregate.FullBook, error)
	UpdateBook()
	DeleteBook()
}

type BookService struct {
	db         *database.Database
	bookRepo   repos.IBookRepository
	authorRepo repos.IAuthorRepository
}

func NewBookService(db *database.Database, bookRepo repos.IBookRepository, authorRepo repos.IAuthorRepository) *BookService {
	return &BookService{
		db:         db,
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

// return entity.Book
func (s *BookService) GetBooks() {

}

func (s *BookService) GetBook() {

}

func (s *BookService) CreateBook(ctx context.Context, agg aggregate.CreateBook) (aggregate.FullBook, error) {
	fullBook := aggregate.FullBook{}
	ctx, f, err := s.db.Session(ctx)
	if err != nil {
		return fullBook, err
	}

	defer func() {
		f(err)
	}()

	book, err := s.bookRepo.CreateBook(ctx, agg.Book) // todo: create book and return id's
	if err != nil {
		return fullBook, err
	}

	authors, err := s.authorRepo.AssignBookByIds(ctx, agg.AuthorIds)
	if err != nil {
		return aggregate.FullBook{}, err
	}
	fullBook.BookEntity = book
	fullBook.Authors = authors

	return fullBook, nil
}

func (s *BookService) UpdateBook() {

}

func (s *BookService) DeleteBook() {

}
