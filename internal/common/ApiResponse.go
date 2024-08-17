package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type responseData struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type responseError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
}

type responseValidation struct {
	StatusCode int      `json:"status_code"`
	Message    []string `json:"message"`
}

func (p *Meta) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}

type Meta struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func ResponseSuccessfully(data interface{}, message string) *responseData {
	return &responseData{
		Data:    data,
		Message: message,
	}
}

func ResponseError(statusCode int, root error, msg string) *responseError {
	return &responseError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
	}
}

func ResponseValidationErrors(statusCode int, err error) *responseValidation {
	errs := err.(validator.ValidationErrors)
	var messages []string

	for _, err := range errs {
		var sb strings.Builder

		sb.WriteString("Validation failed on fieled on field '" + err.Field() + "'")
		sb.WriteString(", condition: " + err.ActualTag())

		if err.Param() != "" {
			sb.WriteString(" { " + err.Param() + " } ")
		}

		if err.Value() != nil && err.Value() != "" {
			sb.WriteString(fmt.Sprintf(", actual: %v", err.Value()))
		}

		messages = append(messages, sb.String())
	}

	return &responseValidation{
		StatusCode: statusCode,
		Message:    messages,
	}
}
