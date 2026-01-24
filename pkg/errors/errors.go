package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ReferenceCode int

const (
	Internal ReferenceCode = iota + 1
	MalformedData
	InvalidData
	Unauthorized
	NotFound
)

type Error struct {
	ReferenceCode ReferenceCode `json:"referenceCode"`
	StatusCode    int           `json:"statusCode"`
	Err           any           `json:"err,omitempty"`
}

func InternalError(err any) Error {
	return Error{
		ReferenceCode: Internal,
		Err:           err,
	}
}

func InternalServerError(err any) Error {
	return Error{
		ReferenceCode: Internal,
		StatusCode:    http.StatusInternalServerError,
		Err:           err,
	}
}

func NotFoundError(err any) Error {
	return Error{
		ReferenceCode: NotFound,
		StatusCode:    http.StatusNotFound,
		Err:           err,
	}
}

func (e Error) Error() string {
	return errToString(e.Err)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func errToString(msg any) string {
	if m, ok := msg.(map[string]any); ok {
		var values []string
		for _, value := range m {
			values = append(values, value.(string))
		}
		return strings.Join(values, ", ")
	}
	return fmt.Sprintf("%v", msg)
}
