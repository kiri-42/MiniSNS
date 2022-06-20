package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

// errorインタフェースをError()を実装
func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

func newHTTPError(code int, msg string) *httpError {
	return &httpError{
		code:    code,
		Key:     http.StatusText(code),
		Message: msg,
	}
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  string
	)

	if he, ok := err.(*httpError); ok {
		code = he.code
		msg = he.Message
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, newHTTPError(code, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
