package responses

// ErrorResponse model
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse returns the json response for error
func NewErrorResponse(error string) *ErrorResponse {
	return &ErrorResponse{Error: error}
}
