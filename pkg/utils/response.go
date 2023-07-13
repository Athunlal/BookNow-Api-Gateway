package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

//In Go, the strings.Split() function is used to split a string into substrings based on a specified delimiter. It returns a slice of substrings.

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(message string, err string, data interface{}) Response {
	splittedstring := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedstring,
		Data:    data,
	}
	return res
}

func SuccessResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
	return res
}

func ResponseJSON(c gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(data)
}

func ExtractError(err error) (string, error) {
	// Check if the error is a gRPC error
	if errStatus, ok := status.FromError(err); ok {
		// Extract the error message from the gRPC error
		errorMessage := errStatus.Message()
		return errorMessage, nil
	} else {
		// Handle non-gRPC errors here
		return "", errors.New("Not a grpc error")
	}
}

func JsonInputValidation(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Success": false,
		"Message": "client-side input validation failed",
		"Error":   "Error in Binding the JSON Data",
	})
}
