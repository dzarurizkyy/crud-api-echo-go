package utils

import "github.com/labstack/echo/v4"

// Struct
type JSONResponse struct {
	Code    int         `json:"code"`;
	Message string      `json:"message"`;
	Data    interface{} `json:"data"`;
};

// Initialize format response body
func SetResponse(c echo.Context, statusCode int, message string, data interface{}) (error) {
	return c.JSON(statusCode, JSONResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	});
};