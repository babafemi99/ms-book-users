package msErrors

import "net/http"

type RestErrors struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequest(message string, err error) *RestErrors {
	return &RestErrors{
		Message: message + "- " + err.Error(),
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundRequestError(message string, err error) *RestErrors {
	return &RestErrors{
		Message: message + "- " + err.Error(),
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErrors {
	return &RestErrors{
		Message: message + "- " + err.Error(),
		Status:  http.StatusInternalServerError,
		Error:   "internal server Error",
	}
}
