package main

import (
	"github.com/Azunyan1111/ad-network/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/dsp", model.Dsp())
	e.GET("/", model.OkDspHD())

	// サーバー起動
	e.Start(":" + os.Getenv("PORT2")) //ポート番号指定してね
}
