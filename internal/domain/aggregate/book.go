package aggregate

import "basic-microservice/internal/domain/entity"

type FullBook struct {
	entity.Book
	Authors    []entity.Author   `json:"authors" gorm:"many2many:author_books;"`
	Categories []entity.Category `json:"categories" gorm:"many2many:book_categories;"`
}

type CreateBook struct {
	Book        entity.Book
	AuthorIds   []int `json:"authorIds"`
	CategoryIds []int `json:"categoryIds"`
}
