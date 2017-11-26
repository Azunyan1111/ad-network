package main

import (
	"github.com/Azunyan1111/ad-network/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	// スタックの起動
	go model.Hoge(model.Stack)

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/imp", model.ImpHD())
	e.GET("/", model.OkHD())

	// サーバー起動
	e.Start(":" + os.Getenv("PORT3")) //ポート番号指定してね
}
