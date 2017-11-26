package model

import (
	"github.com/labstack/echo"
	"net/http"
)

func Ad() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		return c.String(http.StatusOK, "Hello Ad World")
	}
}
