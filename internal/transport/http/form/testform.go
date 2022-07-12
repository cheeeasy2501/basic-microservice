package form

import "errors"

type Forms interface {
	//LoadAndValidate(ctx *gin.Context)
	LoadAndValidate() []error
}

type CreateBookForm struct {
	Title        string
	Descriptions string
}

func (f *CreateBookForm) LoadAndValidate() []error {
	var err []error // struct with errors
	if f.Title == "" {
		err = append(err, errors.New("empty title"))
	}

	if len(err) > 0 {
		return err
	}

	return nil
}
