package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order/errors"
)

func PanicExceptionHandling() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					switch v := err.(type) {
					case *errors.Error:
						v.Log()
						_ = c.JSON(v.StatusCode, v)
					case errors.Error:
						v.Log()
						_ = c.JSON(v.StatusCode, v)
					default:
						errors.Define("Unknown error occurred", 500)
						_ = c.NoContent(http.StatusInternalServerError)
					}
				}
			}()
			return next(c)
		}
	}
}
