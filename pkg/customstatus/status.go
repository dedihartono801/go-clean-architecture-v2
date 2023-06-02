package customstatus

import "net/http"

var (
	ErrUnprocessableEntity = NewStatus(http.StatusUnprocessableEntity, "Unprocessable entity")
	ErrNotFound            = NewStatus(http.StatusNotFound, "Data not found")
	ErrInternalServerError = NewStatus(http.StatusInternalServerError, "Internal Server Error")
	ErrBadRequest          = NewStatus(http.StatusBadRequest, "Bad Request")
	ErrEmailNotFound       = NewStatus(http.StatusNotFound, "Email not found")
	ErrPasswordWrong       = NewStatus(http.StatusNotFound, "Wrong password")
	StatusOk               = NewStatus(http.StatusOK, "Success")
	StatusCreated          = NewStatus(http.StatusCreated, "Success")
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewStatus(code int, message string) *Status {
	return &Status{
		Code:    code,
		Message: message,
	}
}
