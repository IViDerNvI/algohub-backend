// Package core provides the core error handling and response structure for the application.
package core

// Error codes for common HTTP errors.
var (
	ErrJSONFormation = NewErrCode(400, 1000, "JSON format error")
)

type ErrCode struct {
	HTTPCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func NewErrCode(httpcode int, code int, message string) ErrCode {
	return ErrCode{
		HTTPCode: httpcode,
		Code:     code,
		Message:  message,
	}
}

func (e ErrCode) Error() string {
	return e.Message
}
