package form

import "errors"

type CreateBookForm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (f *CreateBookForm) LoadAndValidate() *ErrorResponse {
	errRes := newFormErrorResponse() // struct with errors
	if f.Title == "" {
		errRes.AddError("title", errors.New("title is empty")) // make errors like const
	}

	if len(f.Title) < 5 {
		errRes.AddError("title", errors.New("title must be > 5 "))
	}

	if len(errRes.errs) > 0 {
		return errRes
	}

	return nil
}
