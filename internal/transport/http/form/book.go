package form

const (
	bookTitleNotEmpty = ValidationError("title is empty")
	bookTitleLength   = ValidationError("title must be > 5")
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
		errRes.addError("title", bookTitleNotEmpty) // make errors like const
	}

	if len(f.Title) < 5 {
		errRes.addError("title", bookTitleLength)
	}

	if len(errRes.Errs) > 0 {
		return errRes
	}

	return nil
}
