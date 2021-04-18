package main

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func InternalServerError(err string) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: err,
	}
}

func NotFoundError(err string) Response {
	return Response{
		Status:  http.StatusNotFound,
		Message: err,
	}
}

func BadRequestError(err string) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: err,
	}
}

func CreatedResponse(message string) Response {
	return Response{
		Status: http.StatusCreated,
		Message: message,
	}
}

func DeletedResponse(message string) Response {
	return Response{
		Status: http.StatusOK,
		Message: message,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func InvalidInputError(errs validation.Errors) Response {
	var details []invalidField
	var fields []string

	for field := range errs {
		fields = append(fields, field)
	}

	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return Response{
		Status:  http.StatusBadRequest,
		Message: "Validation Errors on some input fields",
		Details: details,
	}
}
