package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, message string, errors interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// NotFoundResponse sends a not found response
func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: message,
	})
}

// InternalErrorResponse sends an internal server error response
func InternalErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: "Internal server error",
		Error:   err.Error(),
	})
}
