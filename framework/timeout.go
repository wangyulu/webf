package framework

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func TimeoutHandler(handler ControllerHandler, timeout time.Duration) ControllerHandler {
	return func(c *Context) error {
		durationCtx, cancelFunc := context.WithTimeout(c.ctx, timeout*time.Second)
		defer cancelFunc()

		finishChan := make(chan struct{})
		panicChan := make(chan interface{})

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			// 这里是具体的业务逻辑
			// 为什么没有处理 err 的情况 --todo
			handler(c)

			finishChan <- struct{}{}
		}()

		select {
		case <-panicChan:
			c.Json(http.StatusInternalServerError, "panic err")
		case <-finishChan:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.Json(http.StatusInternalServerError, "timeout")
		}

		return nil
	}
}
