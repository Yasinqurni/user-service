package response

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

func NewErrorResponse(message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Error:   true,
		Message: message,
		Err:     err,
	}
}

type ErrorResponseValidate struct {
	Error   bool          `json:"error"`
	Message DescribeError `json:"message"`
	Err     error         `json:"err,omitempty"`
}

type DescribeError struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func NewErrorResponseValidate(describeError DescribeError, err error) *ErrorResponseValidate {
	return &ErrorResponseValidate{
		Error:   true,
		Message: describeError,
		Err:     err,
	}
}
