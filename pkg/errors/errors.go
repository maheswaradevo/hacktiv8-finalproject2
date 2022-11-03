package errors

import (
	"errors"
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/dto"
)

var (
	ErrUnknown            = errors.New("internal server error")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrNotFound           = errors.New("data not found")
	ErrUserExists         = errors.New("email is already taken")
	ErrInvalidResources   = errors.New("resources is empty")
	ErrInvalidCred        = errors.New("password is invalid")
	ErrUnauthorized       = errors.New("user is unauthorized")
	ErrDataNotFound       = errors.New("data not found")
)

func NewErrorData(code int, message string) dto.ErrorData {
	return dto.ErrorData{
		Code:    code,
		Message: message,
	}
}

func GetErrorResponseMetaData(err error) (er dto.ErrorData) {
	er, ok := errorMap[err]
	if !ok {
		er = errorMap[ErrUnknown]
	}
	return
}

var errorMap = map[error]dto.ErrorData{
	ErrUnknown:            NewErrorData(http.StatusInternalServerError, ErrUnknown.Error()),
	ErrInvalidRequestBody: NewErrorData(http.StatusBadRequest, ErrInvalidRequestBody.Error()),
	ErrNotFound:           NewErrorData(http.StatusNotFound, ErrNotFound.Error()),
	ErrUserExists:         NewErrorData(http.StatusBadRequest, ErrUserExists.Error()),
	ErrInvalidResources:   NewErrorData(http.StatusNotFound, ErrInvalidResources.Error()),
	ErrInvalidCred:        NewErrorData(http.StatusBadRequest, ErrInvalidCred.Error()),
	ErrUnauthorized:       NewErrorData(http.StatusUnauthorized, ErrUnauthorized.Error()),
	ErrDataNotFound:       NewErrorData(http.StatusNotFound, ErrDataNotFound.Error()),
}
