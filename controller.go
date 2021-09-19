package main

import (
	"context"
	"fmt"
	"geek/webf/framework"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	durationCtx, cancelFunc := context.WithTimeout(ctx.BaseContext(), time.Second)

	defer cancelFunc()

	panicChan := make(chan interface{})
	finish := make(chan struct{})

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// 具体的业务逻辑
		time.Sleep(10 * time.Second)

		ctx.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "timeout")

		ctx.SetHasTimeout()
	}

	return nil
}
