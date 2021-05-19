package utils

type ErrorResponse struct {
	Error string `json:"error"`
}
type ValidationErrorResponse struct {
	Errors []*ValidationError `json:"errors"`
}

func NewErrorResponse(error error) ErrorResponse {
	return ErrorResponse{error.Error()}
}

func NewValidationErrorResponse(errors []*ValidationError) ValidationErrorResponse {
	return ValidationErrorResponse{errors}
}
