package model

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"github.com/google/uuid"
)

func Dsp() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		dspRes := dspResponse{
			AdId:     c.QueryParam("adid"),
			DspId:    1,
			ImpId:    uuid.New().String(),
			Price:    rand.Intn(200),
			Creative: `<img src="http://dic.nicovideo.jp/oekaki/185531.png">`,
		}
		return c.JSON(http.StatusOK, dspRes)
	}
}

func OkDspHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fast(c)
		return c.JSON(http.StatusOK, okResponse{Status: "Hello Dsp Server"})
	}
}
