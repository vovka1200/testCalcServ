package main

import (
	"encoding/json"
	"errors"
	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Answer struct {
	Success bool   `json:"success"`
	ErrCode string `json:"errcode"`
	Value   int    `json:"value"`
}

func main() {
	// Запуск сервера
	log.Info("Запуск...")
	// Роутер
	r := router.New()
	r.GET("/add", servAdd)
	r.GET("/sub", servSub)
	r.GET("/mul", servMul)
	r.GET("/div", servDiv)

	if err := fasthttp.ListenAndServe("0.0.0.0:8080", r.Handler); err != nil {
		log.Fatal(err)
	}
}

func getParams(ctx *fasthttp.RequestCtx) (int, int, error) {
	if a, err := strconv.ParseInt(string(ctx.FormValue("a")), 10, 32); err == nil {
		if b, err := strconv.ParseInt(string(ctx.FormValue("b")), 10, 32); err == nil {
			return int(a), int(b), nil
		} else {
			return 0, 0, err
		}
	} else {
		return 0, 0, err
	}
}

func servAdd(ctx *fasthttp.RequestCtx) {
	if a, b, err := getParams(ctx); err == nil {
		returnResult(ctx, a+b)
	} else {
		returnError(ctx, err)
	}
}

func servSub(ctx *fasthttp.RequestCtx) {
	if a, b, err := getParams(ctx); err == nil {
		returnResult(ctx, a-b)
	} else {
		returnError(ctx, err)
	}
}

func servMul(ctx *fasthttp.RequestCtx) {
	if a, b, err := getParams(ctx); err == nil {
		returnResult(ctx, a*b)
	} else {
		returnError(ctx, err)
	}
}

func servDiv(ctx *fasthttp.RequestCtx) {
	if a, b, err := getParams(ctx); err == nil {
		if b != 0 {
			returnResult(ctx, a/b)
		} else {
			returnError(ctx, errors.New("деление на ноль"))
		}
	} else {
		returnError(ctx, err)
	}
}

func returnResult(ctx *fasthttp.RequestCtx, a int) {
	if buf, err := json.Marshal(Answer{
		Success: true,
		ErrCode: "",
		Value:   a,
	}); err == nil {
		ctx.SetBody(buf)
	} else {
		returnError(ctx, err)
	}
}

func returnError(ctx *fasthttp.RequestCtx, err error) {
	if buf, err := json.Marshal(Answer{
		Success: false,
		ErrCode: err.Error(),
		Value:   0,
	}); err == nil {
		ctx.SetBody(buf)
	}
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
}
