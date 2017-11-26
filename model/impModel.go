package model

import (
	"github.com/labstack/echo"
	"net/http"
)

var Stack = make(chan imp)

func ImpHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		var impRequest imp
		impRequest.impId = c.Request().Form.Get("impId")
		impRequest.dspId = c.Request().Form.Get("dspId")
		Stack <- impRequest
		return c.JSON(http.StatusNoContent, impResponse{Status:"OK"})
	}
}

// スタック（チャンネルの中に構造体が入ると処理が始まる）
func Hoge(stack chan imp){
	for out := range stack{
		// 記録処理
	}
}