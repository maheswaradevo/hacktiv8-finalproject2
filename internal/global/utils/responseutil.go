package utils

import (
	"net/http"

	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
)

func NewSuccessResponseWriter(rw http.ResponseWriter, code int, status string, data interface{}) dto.BaseResponse {
	return BaseResponseWriter(rw, code, status, nil, data)
}

func NewErrorResponse(rw http.ResponseWriter, err error) dto.BaseResponse {
	errMap := errors.GetErrorResponseMetaData(err)
	return BaseResponseWriter(rw, errMap.Code, "", &dto.ErrorData{Code: errMap.Code, Message: errMap.Message}, nil)
}

func BaseResponseWriter(rw http.ResponseWriter, code int, status string, er *dto.ErrorData, data interface{}) dto.BaseResponse {
	res := dto.BaseResponse{
		Status: status,
		Data:   data,
		Error:  er,
	}
	return res
}
