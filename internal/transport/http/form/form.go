package form

type IForm interface {
	LoadAndValidate() []error
}

type ValidationError []error

type ErrorResponse struct {
	message string
	errs    map[string]ValidationError
}

func (r *ErrorResponse) AddError(fieldName string, err error) {
	if _, ok := r.errs[fieldName]; ok == false {
		r.errs[fieldName] = append(r.errs[fieldName], err)
		return
	}

	//r.errs[fieldName] = []
}

func newFormErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		message: "Validation error",
		errs:    make(map[string]ValidationError, 0),
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
