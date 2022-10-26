package errors

import (
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/dto"
)

var (
	ErrUnknown            error
	ErrInvalidRequestBody error
	ErrNotFound           error
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
	ErrUnknown:            NewErrorData(http.StatusInternalServerError, "Internal Server Error"),
	ErrInvalidRequestBody: NewErrorData(http.StatusBadRequest, "Invalid Request Body"),
	ErrNotFound:           NewErrorData(http.StatusNotFound, "Data Not Found"),
}
