package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter

	ctx context.Context

	hasTimeout bool

	writerMux *sync.Mutex

	handler ControllerHandler
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:        r,
		responseWriter: w,

		ctx: r.Context(),

		writerMux: &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHandler(handler ControllerHandler) {
	ctx.handler = handler
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// implement context.context
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryInt(key string, def int) int {
	if val, ok := ctx.QueryAll()[key]; ok {
		if length := len(val); length > 0 {
			res, err := strconv.Atoi(val[length-1])
			if err != nil {
				return def
			}

			return res
		}
	}

	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	if val, ok := ctx.QueryAll()[key]; ok {
		if length := len(val); length > 0 {
			return val[length-1]
		}
	}

	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	if val, ok := ctx.QueryAll()[key]; ok {
		return val
	}

	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}

	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	if val, ok := ctx.FormAll()[key]; ok {
		length := len(val)
		if length > 0 {
			res, err := strconv.Atoi(val[length-1])
			if err != nil {
				return def
			}

			return res
		}
	}

	return def
}

func (ctx *Context) FormString(key string, def string) string {
	if val, ok := ctx.FormAll()[key]; ok {
		if length := len(val); length > 0 {
			return val[length-1]
		}
	}

	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	if val, ok := ctx.FormAll()[key]; ok {
		return val
	}

	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}

	return map[string][]string{}
}

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}

	ctx.responseWriter.Header().Set("Content-type", "application/json")
	ctx.responseWriter.WriteHeader(status)

	bytes, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// 这里不用处理 err 的情况吗 --todo
	ctx.responseWriter.Write(bytes)

	return nil
}
