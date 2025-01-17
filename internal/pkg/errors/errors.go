package errors

import (
	"encoding/json"
	"fmt"
	"strings"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rotisserie/eris"
)

// Error represents a single validation error
type Error struct {
	Field    string    `json:"field"`
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewValidationError(field string, code ErrorCode, message string) *Error {
	return &Error{
		Field:   field,
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Message
}

func IsError(err error) bool {
	var validationErr *Error
	var multiErr *MultiError
	return eris.As(err, &validationErr) || eris.As(err, &multiErr)
}

type MultiError struct {
	Errors []*Error `json:"errors"`
}

func NewMultiError() *MultiError {
	return &MultiError{
		Errors: make([]*Error, 0),
	}
}

// Add adds a new validation error to the collection
func (m *MultiError) Add(field string, code ErrorCode, message string) {
	m.Errors = append(m.Errors, &Error{
		Field:   field,
		Code:    code,
		Message: message,
	})
}

// AddError adds an existing Error to the collection
func (m *MultiError) AddError(err *Error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

// HasErrors returns true if there are any validation errors
func (m *MultiError) HasErrors() bool {
	return len(m.Errors) > 0
}

// Error implements the error interface
func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}

	var messages []string
	for _, err := range m.Errors {
		messages = append(messages, err.Error())
	}

	return fmt.Sprintf("validation failed:\n- %s", strings.Join(messages, "\n- "))
}

func (m *MultiError) MarshalJSON() ([]byte, error) {
	if m == nil || len(m.Errors) == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(struct {
		Errors []*Error `json:"errors"`
	}{
		Errors: m.Errors,
	})
}

func IsMultiError(err error) bool {
	var multiErr *MultiError
	return eris.As(err, &multiErr)
}

type BusinessError struct {
	Code     ErrorCode         `json:"code"`
	Message  string            `json:"message"`
	Details  string            `json:"details,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
	Internal error             `json:"-"`
}

func NewBusinessError(message string) *BusinessError {
	return &BusinessError{
		Code:    ErrBusinessLogic,
		Message: message,
	}
}

func (e *BusinessError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

func (e *BusinessError) WithParam(key, value string) *BusinessError {
	if e.Params == nil {
		e.Params = make(map[string]string)
	}
	e.Params[key] = value
	return e
}

func (e *BusinessError) WithInternal(err error) *BusinessError {
	e.Internal = err
	return e
}

func IsBusinessError(err error) bool {
	var businessErr *BusinessError
	return eris.As(err, &businessErr)
}

type DatabaseError struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{
		Code:    ErrSystemError,
		Message: message,
	}
}

func (e *DatabaseError) Error() string {
	return e.Message
}

func IsDatabaseError(err error) bool {
	var databaseErr *DatabaseError
	return eris.As(err, &databaseErr)
}

func (e *DatabaseError) WithInternal(err error) *DatabaseError {
	e.Internal = err
	return e
}

type AuthenticationError struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		Code:    ErrUnauthorized,
		Message: message,
	}
}

func IsAuthenticationError(err error) bool {
	var authenticationErr *AuthenticationError
	return eris.As(err, &authenticationErr)
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func (e *AuthenticationError) WithInternal(err error) *AuthenticationError {
	e.Internal = err
	return e
}

type AuthorizationError struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{
		Code:    ErrForbidden,
		Message: message,
	}
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

func IsAuthorizationError(err error) bool {
	var authorizationErr *AuthorizationError
	return eris.As(err, &authorizationErr)
}

func (e *AuthorizationError) WithInternal(err error) *AuthorizationError {
	e.Internal = err
	return e
}

type NotFoundError struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Code:    ErrNotFound,
		Message: message,
	}
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func IsNotFoundError(err error) bool {
	var notFoundErr *NotFoundError
	return eris.As(err, &notFoundErr)
}

type RateLimitError struct {
	Field    string    `json:"field,omitempty"`
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Internal error     `json:"-"`
}

func NewRateLimitError(field string, message string) *RateLimitError {
	return &RateLimitError{
		Field:   field,
		Code:    ErrTooManyRequests,
		Message: message,
	}
}

func IsRateLimitError(err error) bool {
	var rateLimitErr *RateLimitError
	return eris.As(err, &rateLimitErr)
}

func (e *RateLimitError) Error() string {
	return e.Message
}

func (e *RateLimitError) WithInternal(err error) *RateLimitError {
	e.Internal = err
	return e
}

func inferErrorCode(err error) ErrorCode {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "required"):
		return ErrRequired
	case strings.Contains(msg, "length"):
		return ErrInvalidLength
	case strings.Contains(msg, "format"):
		return ErrInvalidFormat
	case strings.Contains(msg, "match"):
		return ErrInvalidFormat
	case strings.Contains(msg, "rate limit"):
		return ErrTooManyRequests
	default:
		return ErrInvalid
	}
}

func FromValidationErrors(valErrors val.Errors, multiErr *MultiError, prefix string) {
	for field, err := range valErrors {
		fieldName := field
		if prefix != "" {
			fieldName = fmt.Sprintf("%s.%s", prefix, field)
		}

		multiErr.AddError(&Error{
			Field:   fieldName,
			Code:    inferErrorCode(err),
			Message: err.Error(),
		})
	}
}