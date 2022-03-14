//Package errCode - define all common error in this server
package errCode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code   int      `json:"Code"`
	Msg    string   `json:"Msg"`
	Detail []string `json:"Detail"`
}

var codes = map[int]string{} //errCode : Msg
func NewError(code int, msg string) *Error {
	//if exist ,panic
	if _, ok := codes[code]; ok {
		//found
		panic(fmt.Sprintf("Error Code %d is already exist", code))
	}

	codes[code] = msg
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintf("")
}

func (err *Error) GetCode() int {
	return err.Code
}

func (err *Error) GetMsg() string {
	return err.Msg
}

func (err *Error) GetDetail() []string {
	return err.Detail
}

func (err *Error) WithDetail(detail ...string) *Error {
	newErr := *err
	newErr.Detail = []string{}

	//append passing Detail to error Detail
	for _, data := range detail {
		newErr.Detail = append(newErr.Detail, data)
	}

	return &newErr
}

func (err *Error) StatusCode() int {

	switch err.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case PermanentlyRedirect.GetCode():
		return http.StatusMovedPermanently
	case ServerError.GetCode():
		return http.StatusInternalServerError
	case InvalidParams.GetCode():
		return http.StatusBadRequest
	case TooManyRequest.GetCode():
		return http.StatusTooManyRequests
	case NotFound.GetCode():
		return http.StatusNotFound
	case ErrorCreateShortenURL.GetCode():
		return http.StatusInternalServerError
	case ErrorGetURL.GetCode():
		fallthrough
	case ErrorUrlCodeExpired.GetCode():
		break
	}

	return http.StatusInternalServerError
}
