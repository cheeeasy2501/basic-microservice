package form

import (
	"errors"
)

type CreateBookForm struct {
	Isbn            string `json:"isbn"`
	Status          string `json:"status"`
	Title           string `json:"title"`
	LanguageLevelId int    `json:"languageLevelId"`
	Description     string `json:"description"`
	Link            string `json:"link"`
	CoverPath       string `json:"coverPath"`
	AuthorIds       []int  `json:"authorIds"`
}

func (f *CreateBookForm) LoadAndValidate() *ErrorResponse {
	errRes := newFormErrorResponse() // struct with errors
	if f.Title == "" {
		errRes.AddError("title", errors.New("title is empty")) // make errors like const
	}

	if len(f.Title) < 5 {
		errRes.AddError("title", errors.New("title must be > 5"))
	}

	if len(errRes.Errs) > 0 {
		return errRes
	}

	return nil
}
