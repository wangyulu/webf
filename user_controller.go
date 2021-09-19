package main

import (
	"geek/webf/framework"
	"net/http"
)

func UserLoginController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, UserLoginController")
	return nil
}
