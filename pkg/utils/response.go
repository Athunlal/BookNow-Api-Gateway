package utils

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"erro
	rs,omitempty"`
	Data interface{} `json:"data,omitempty"`
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

func SeatValidation(data string) error {
	pattern := regexp.MustCompile(`A[1-9]|A1[0-9]|A20|B[1-9]|B1[0-9]|B20|S[1-9]|S1[0-9]|S20`)
	if !pattern.MatchString(data) {
		return fmt.Errorf("Compartment name should be A1 to A20 , B1 t0 B2 and S1 to S20")
	}
	return nil
}

func ConvertStringToTimestamp(input string) (primitive.Timestamp, error) {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return primitive.Timestamp{}, err
	}

	seconds := parsedTime.Unix()
	nanos := uint32(parsedTime.Nanosecond()) // Convert int32 to uint32

	return primitive.Timestamp{T: uint32(seconds), I: nanos}, nil
}

func ConvertTimestampToString(timestamp primitive.Timestamp) string {
	// Convert primitive.Timestamp to time.Time
	seconds := int64(timestamp.T)
	nanos := int64(timestamp.I)
	t := time.Unix(seconds, nanos)

	// Format time.Time as a string (change format as needed)
	return t.Format(time.RFC3339)
}

func TimeToTimestamp(normalTime string) (primitive.Timestamp, error) {
	parts := strings.Split(normalTime, ":")
	if len(parts) != 2 {
		return primitive.Timestamp{}, fmt.Errorf("invalid time format: %s", normalTime)
	}

	hour := parts[0]
	minute := parts[1]

	hourInt := 0
	minuteInt := 0

	_, err := fmt.Sscanf(hour, "%d", &hourInt)
	if err != nil {
		return primitive.Timestamp{}, err
	}

	_, err = fmt.Sscanf(minute, "%d", &minuteInt)
	if err != nil {
		return primitive.Timestamp{}, err
	}

	now := time.Now()
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), hourInt, minuteInt, 0, 0, now.Location())

	seconds := targetTime.Unix()
	return primitive.Timestamp{
		T: uint32(seconds),
		I: 0,
	}, nil
}

func BindUpdateData(ctx *gin.Context) (domain.Train, error) {
	updateData := domain.Train{}
	err := ctx.Bind(&updateData)
	return updateData, err
}
