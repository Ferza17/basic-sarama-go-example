package util

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ferza17/kafka-basic/consumer/exception"
)

type ResponseRequest struct {
	Exception  *exception.Exception
	Data       interface{}
	StatusCode int
}

// Response
func Response(c echo.Context, r ResponseRequest) error {
	if r.Exception != nil {
		return r.responseError(c)
	}
	return r.responseSuccess(c)
}

// responseError : handling error response
func (r *ResponseRequest) responseError(c echo.Context) error {
	response := make(map[string]interface{})
	response["ErrorCode"] = r.Exception.ErrorCode
	response["Message"] = r.Exception.Message
	return c.JSON(r.Exception.StatusCode, response)
}

// responseSuccess : handling success response
func (r *ResponseRequest) responseSuccess(c echo.Context) error {
	var (
		response   = make(map[string]interface{})
		statusCode = http.StatusOK
	)

	if r.StatusCode != statusCode {
		statusCode = r.StatusCode
	}

	response["Message"] = "Success"

	response["Data"] = r.Data

	return c.JSON(statusCode, response)
}
