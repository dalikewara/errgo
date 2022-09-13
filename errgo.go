package errgo

import "errors"

type ErrGo interface {
	GetError() error
	GetCode() string
	GetMessage() string
	GetHttpStatus() int
}

type errgo struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
	Error      error  `json:"error"`
}

// New generates new errgo that implements ErrGo.
func New(code, message string) ErrGo {
	return &errgo{
		Code:    code,
		Message: message,
		Error:   generateError(code, message, "||"),
	}
}

// NewWithHttpStatus generates new errgo with http status that implements ErrGo.
func NewWithHttpStatus(code, message string, httpStatus int) ErrGo {
	return &errgo{
		Code:       code,
		Message:    message,
		HttpStatus: httpStatus,
		Error:      generateError(code, message, "||"),
	}
}

// GetError gets errgo error.
func (e *errgo) GetError() error {
	return e.Error
}

// GetCode gets errgo code.
func (e *errgo) GetCode() string {
	return e.Code
}

// GetMessage gets errgo message.
func (e *errgo) GetMessage() string {
	return e.Message
}

// GetHttpStatus gets errgo http status.
func (e *errgo) GetHttpStatus() int {
	return e.HttpStatus
}

// generateError generates new error.
func generateError(code, message, separator string) error {
	return errors.New(code + separator + message)
}
