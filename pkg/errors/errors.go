package errors

import (
	"fmt"
	"gin_test/pkg/app"
	"gin_test/pkg/logging"

	"github.com/pkg/errors"
)

// ErrorType is the type of an error
type ErrorType uint

const (
	// NoType error
	NoType ErrorType = iota
	// BadRequest error
	BadRequest
	// NotFound error
	NotFound
)

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	Code    int
	Field   string
	Message string
}

// New creates a new customError
func (errorType ErrorType) New(msg string) error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// New creates a new customError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New creates a no type error
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(code int, msg, field string, args ...interface{}) error {
	// 没有自定义错误内容取统一定义错误内容
	if msg == "" {
		msg = app.GetMsg(code)
	}

	err := New(msg)
	errContext := errorContext{
		Code:    code,
		Field:   field,
		Message: msg,
	}
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       errContext,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError, context: errContext}
}

// NewErrf an error with format string
func NewErrf(err error, code int, msg, field string, args ...interface{}) error {
	// 没有自定义错误内容取统一定义错误内容
	if msg == "" {
		msg = app.GetMsg(code)
	}

	errContext := errorContext{
		Code:    code,
		Field:   field,
		Message: msg,
	}
	wrappedError := Wrapf(err, msg, args...)

	if customErr, ok := err.(customError); ok {
		customErr := customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       errContext,
		}

		return customErr
	}
	// 记录错误日志
	logging.LogErrorWithFields(wrappedError, logging.Fields{})
	return customError{errorType: NoType, originalError: wrappedError, context: errContext}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, code int, message, field string) error {
	context := errorContext{Code: code, Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, context: context}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) errorContext {
	emptyContext := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.context != emptyContext {
		return customErr.context
	}

	return emptyContext
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
