package main

import (
	"geek/webf/framework"
	"geek/webf/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery(), middleware.Cost())

	registerRouter(core)

	server := &http.Server{
		Addr: ":8001",

		Handler: core,
	}

	server.ListenAndServe()
}
