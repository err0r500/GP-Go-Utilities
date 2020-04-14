package utilities

import (
	"strconv"
)

// HTTPError implement ClientError
type HTTPError struct {
	Cause  error  `json:"error"`
	Detail string `json:"detail"`
	Status int    `json:"code"`
}

func (e *HTTPError) Error() string {
	return "[" + strconv.Itoa(e.Status) + "] " + e.Detail + " : " + e.Cause.Error()
}

func NewHTTPError(err error, status int, detail string) error {
	if err == nil {
		return nil
	}
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}

}
