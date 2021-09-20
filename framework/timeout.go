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
			// 执行具体的业务逻辑
			time.Sleep(10 * time.Second)

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
