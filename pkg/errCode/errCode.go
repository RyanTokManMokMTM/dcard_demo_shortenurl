//Package errCode - define all common error in this server
package errCode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code   int      `json:"code"`
	msg    string   `json:"msg"`
	detail []string `json:"detail"`
}

var codes = map[int]string{} //errCode : msg
func NewError(code int, msg string) *Error {
	//if exist ,panic
	if _, ok := codes[code]; ok {
		//found
		panic(fmt.Sprintf("Error code %d is already exist", code))
	}

	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintf("")
}

func (err *Error) Code() int {
	return err.code
}

func (err *Error) Msg() string {
	return err.msg
}

func (err *Error) Detail() []string {
	return err.detail
}

func (err *Error) WithDetail(detail ...string) *Error {
	newErr := *err
	newErr.detail = []string{}

	//append passing detail to error detail
	for _, data := range detail {
		newErr.detail = append(newErr.detail, data)
	}

	return &newErr
}

func (err *Error) StatusCode() int {
	switch err {
	case Success:
		return http.StatusOK
	case ServerError:
		return http.StatusInternalServerError
	case InvalidParams:
		return http.StatusBadRequest
	case TooManyRequest:
		return http.StatusTooManyRequests
	case NotFound:
		return http.StatusNotFound
	case ErrorCreateShortenURL:
		fallthrough
	case ErrorGetURL:
		fallthrough
	case ErrorUrlCodeExpired:
	}

	return http.StatusInternalServerError
}
