package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree

	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Tree{}

	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{
		router: router,

		middlewares: []ControllerHandler{},
	}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serverHTTP")

	ctx := NewContext(w, r)

	router := c.FindRouteByRequest(r)
	if router == nil {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}

	log.Println("core.router")

	ctx.SetHandler(router...)

	if err := ctx.Next(); err != nil {
		ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(strings.ToUpper(url), allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(strings.ToUpper(url), allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(strings.ToUpper(url), allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(strings.ToUpper(url), allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	uri := strings.ToUpper(request.URL.Path)
	method := strings.ToUpper(request.Method)

	if routers, ok := c.router[method]; ok {
		return routers.FindHandler(uri)
	}

	return nil
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
