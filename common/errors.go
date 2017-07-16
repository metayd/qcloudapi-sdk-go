package common

import (
	"fmt"
)

const (
	NoErr         = 0
	NoErrCodeDesc = "Success"

	ErrQCloudGoClient = 9999
)

type Error struct {
	ErrorResponse ErrorResponse
	StatusCode    int
}

type ErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (err Error) Error() string {
	return fmt.Sprintf(
		"QCloud API Error: HTTP Status Code: %d, Response Code: %d, Message: %s",
		err.StatusCode,
		err.ErrorResponse.Code,
		err.ErrorResponse.Message,
	)
}

func makeClientError(err error) Error {
	return Error{
		ErrorResponse: ErrorResponse{
			Code:    ErrQCloudGoClient,
			Message: err.Error(),
		},
		StatusCode: -1,
	}
}
