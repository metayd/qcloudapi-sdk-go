package clb

import "fmt"

const (
	NoErr         = 0
	NoErrCodeDesc = "Success"

	ErrParamInvalid          = 4000
	ErrAuthFailure           = 4100
	ErrRequestExpired        = 4200
	ErrForbidden             = 4300
	ErrQuotaExceeded         = 4400
	ErrReplayAttack          = 4500
	ErrUnsupportedProtocol   = 4600
	ErrResourceDoesNotExists = 5000
	ErrOperationFailure      = 5100
	ErrPurchaseFailure       = 5200
	ErrOutOfMoney            = 5300
	ErrPartialSuccess        = 5400
	ErrLackOfQualifications  = 5500
	ErrInternalServerError   = 6000
	ErrUnsupportedVersion    = 6100
	ErrEndpointUnavailable   = 6200

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
