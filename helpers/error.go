package helpers

import (
	"strconv"
)

// HTTPError implement ClientError
type HTTPError struct {
	Cause  string `json:"error"`
	Detail string `json:"message"`
	Status int    `json:"code"`
}

func (e *HTTPError) Error() string {
	return "[" + strconv.Itoa(e.Status) + "] " + e.Detail + " : " + e.Cause
}

func NewHTTPError(err error, status int, detail string) error {
	if err == nil {
		return nil
	}
	return &HTTPError{
		Cause:  err.Error(),
		Detail: detail,
		Status: status,
	}

}
