package main

import (
	"geek/webf/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()

	registerRouter(core)

	server := &http.Server{
		Addr: ":8001",

		Handler: core,
	}

	server.ListenAndServe()
}
