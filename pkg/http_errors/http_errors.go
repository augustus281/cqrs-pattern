package httpErrors

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/augustus281/cqrs-pattern/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

const (
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrRequestTimeout      = "Request Timeout"
	ErrInvalidEmail        = "Invalid email"
	ErrInvalidPassword     = "Invalid password"
	ErrInvalidField        = "Invalid field"
	ErrInternalServerError = "Internal Server Error"
)

var (
	BadRequest          = errors.New("Bad request")
	WrongCredentials    = errors.New("Wrong Credentials")
	NotFound            = errors.New("Not Found")
	Unauthorized        = errors.New("Unauthorized")
	Forbidden           = errors.New("Forbidden")
	InternalServerError = errors.New("Internal Server Error")
)

// RestErr Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	ErrBody() RestError
}

// RestError Rest error struct
type RestError struct {
	ErrStatus  int         `json:"status,omitempty"`
	ErrError   string      `json:"error,omitempty"`
	ErrMessage interface{} `json:"message,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
}

// ErrBody Error body
func (e RestError) ErrBody() RestError {
	return e
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrMessage)
}

// Status Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrMessage
}

// NewRestError New Rest Error
func NewRestError(status int, err string, causes interface{}, debug bool) RestErr {
	restError := RestError{
		ErrStatus: status,
		ErrError:  err,
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	return restError
}

// NewRestErrorWithMessage New Rest Error With Message
func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus:  status,
		ErrError:   err,
		ErrMessage: causes,
		Timestamp:  time.Now().UTC(),
	}
}

// NewRestErrorFromBytes New Rest Error From Bytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// NewBadRequestError New Bad Request Error
func NewBadRequestError(ctx *gin.Context, causes interface{}, debug bool) {
	restError := RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  BadRequest.Error(),
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	ctx.JSON(http.StatusBadRequest, restError)
	ctx.Abort()
}

// NewNotFoundError New Not Found Error
func NewNotFoundError(ctx *gin.Context, causes interface{}, debug bool) {
	restError := RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  NotFound.Error(),
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	ctx.JSON(http.StatusNotFound, restError)
	ctx.Abort()
}

// NewUnauthorizedError New Unauthorized Error
func NewUnauthorizedError(ctx *gin.Context, causes interface{}, debug bool) {
	restError := RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  Unauthorized.Error(),
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	ctx.JSON(http.StatusUnauthorized, restError)
	ctx.Abort()
}

// NewForbiddenError New Forbidden Error
func NewForbiddenError(ctx *gin.Context, causes interface{}, debug bool) {
	restError := RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  Forbidden.Error(),
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	ctx.JSON(http.StatusForbidden, restError)
	ctx.Abort()
}

// NewInternalServerError New Internal Server Error
func NewInternalServerError(ctx *gin.Context, causes interface{}, debug bool) {

	restError := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
		Timestamp: time.Now().UTC(),
	}
	if debug {
		restError.ErrMessage = causes
	}
	ctx.JSON(http.StatusInternalServerError, restError)
	ctx.Abort()
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error, debug bool) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, ErrNotFound, err.Error(), debug)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout, err.Error(), debug)
	case errors.Is(err, Unauthorized):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case errors.Is(err, WrongCredentials):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.SQLState):
		return parseSqlErrors(err, debug)
	case strings.Contains(strings.ToLower(err.Error()), "field validation"):
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return NewRestError(http.StatusBadRequest, ErrBadRequest, validationErrors.Error(), debug)
		}
		return parseValidatorError(err, debug)
	case strings.Contains(strings.ToLower(err.Error()), "required header"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Base64):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Unmarshal):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Uuid):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Cookie):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Token):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), constants.Bcrypt):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), "no documents in result"):
		return NewRestError(http.StatusNotFound, ErrNotFound, err.Error(), debug)
	default:
		if restErr, ok := err.(*RestError); ok {
			return restErr
		}
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError, errors.Cause(err).Error(), debug)
	}
}

func parseSqlErrors(err error, debug bool) RestErr {
	return NewRestError(http.StatusBadRequest, ErrBadRequest, err, debug)
}

func parseValidatorError(err error, debug bool) RestErr {
	if strings.Contains(err.Error(), "Password") {
		return NewRestError(http.StatusBadRequest, ErrInvalidPassword, err, debug)
	}

	if strings.Contains(err.Error(), "Email") {
		return NewRestError(http.StatusBadRequest, ErrInvalidEmail, err, debug)
	}

	return NewRestError(http.StatusBadRequest, ErrInvalidField, err, debug)
}

// ErrorResponse Error response
func ErrorResponse(err error, debug bool) (int, interface{}) {
	return ParseErrors(err, debug).Status(), ParseErrors(err, debug)
}

// ErrorCtxResponse Error response object and status code
func ErrorCtxResponse(ctx *gin.Context, err error, debug bool) {
	restErr := ParseErrors(err, debug)
	ctx.JSON(restErr.Status(), restErr)
	ctx.Abort()
}
