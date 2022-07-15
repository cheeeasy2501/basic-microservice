package aggregate

import "basic-microservice/internal/domain/entity"

type FullBook struct {
	entity.BookEntity
	Authors    []entity.AuthorEntity `json:"authors" gorm:"many2many:author_books;"`
	Categories []entity.Category     `json:"categories" gorm:"many2many:book_categories;"`
}

type CreateBook struct {
	Book        entity.BookEntity
	AuthorIds   []int `json:"authorsIds"`
	CategoryIds []int `json:"categoryIds"`
}
