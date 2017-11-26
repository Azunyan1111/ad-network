package model

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

var Stack = make(chan imp)

func ImpHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		var impRequest imp
		impRequest.impId = c.QueryParam("impid")
		impRequest.dspId = c.QueryParam("dspid")
		impRequest.price = c.QueryParam("price")
		if impRequest.impId == "" || impRequest.dspId == "" || impRequest.price == "" {
			return c.JSON(http.StatusBadRequest, okResponse{Status: "NG"})
		}
		Stack <- impRequest
		return c.JSON(http.StatusOK, okResponse{Status: "OK"})
	}
}

// スタック（チャンネルの中に構造体が入ると処理が始まる）（重い処理が最適）
func Hoge(stack chan imp) {
	for out := range stack {
		log.Println(out)
	}
}

func OkImpHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		return c.JSON(http.StatusOK, okResponse{Status: "Hello Imp Server"})
	}
}
