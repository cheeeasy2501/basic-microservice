package form

type IForm interface {
	LoadAndValidate() []error
}

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

//type ValidationErrors []error // todo: not return text in response
type ValidationErrors []ValidationError

type ErrorResponse struct {
	Message string                      `json:"message"`
	Errs    map[string]ValidationErrors `json:"errors"`
}

func (r *ErrorResponse) AddError(fieldName string, err error) {
	r.Errs[fieldName] = append(r.Errs[fieldName], ValidationError(err.Error()))
	return
}

func newFormErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Message: "Validation error",
		Errs:    make(map[string]ValidationErrors, 0),
	}
}

/// Validation
//{
//	message: "Validation Error",
//	errors: {
//	password: [
//		1: "Invalid password",
//		2: "Please,  use @ letter"
//	],
//	login: [
//		1: "abc",
//		2: "def"
//	],
//	}
//}
