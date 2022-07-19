package form

type IForm interface {
	LoadAndValidate() []error
}

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

type ValidationErrors []ValidationError

type ErrorResponse struct {
	Message string                      `json:"message"`
	Errs    map[string]ValidationErrors `json:"errors"`
}

func (r *ErrorResponse) addError(fieldName string, err ValidationError) {
	r.Errs[fieldName] = append(r.Errs[fieldName], err)
	return
}

func newFormErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Message: "Validation error",
		Errs:    make(map[string]ValidationErrors, 0),
	}
}
