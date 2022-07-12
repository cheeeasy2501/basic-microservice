package service

type IBookService interface {
	GetBooks()
	GetBook()
	CreateBook()
	UpdateBook()
	DeleteBook()
}

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// return entity.Book
func (s *BookService) GetBooks() {

}

func (s *BookService) GetBook() {

}
func (s *BookService) CreateBook() {

}
func (s *BookService) UpdateBook() {

}
func (s *BookService) DeleteBook() {

}
