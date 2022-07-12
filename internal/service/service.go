package service

type Services struct {
	bookService IBookService
}

func NewServices() *Services {
	return &Services{
		bookService: NewBookService(),
	}
}
