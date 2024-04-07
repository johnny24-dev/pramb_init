package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
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

func FailureJson(ctx *gin.Context, statusCode int, booleanValue bool, message string, err string) {
	ctx.JSON(statusCode, gin.H{
		"Success": booleanValue,
		"Message": message,
		"Error":   err,
	})
}
