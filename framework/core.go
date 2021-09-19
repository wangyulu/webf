package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serverHTTP")

	ctx := NewContext(w, r)

	router := c.FindRouteByRequest(r)
	if router == nil {
		return
	}

	log.Println("core.router")

	ctx.SetHandler(router)

	// 这里为什么没有处理 error todo
	router(ctx)
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// todo
	return c.router["foo"]
}
