package middleware

import (
	"context"
	"fmt"
	"geek/webf/framework"
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		durationCtx, cancelFunc := context.WithTimeout(c.BaseContext(), timeout*time.Second)
		defer cancelFunc()

		panicChan := make(chan interface{})
		finishChan := make(chan struct{})

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}

				// 这里不用去处理 err 的情况吗 --todo
				c.Next()

				finishChan <- struct{}{}
			}()
		}()

		select {
		case <-panicChan:
			c.Json(http.StatusInternalServerError, "inner err")
		case <-finishChan:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.Json(http.StatusInternalServerError, "timeout")
		}

		return nil
	}
}
