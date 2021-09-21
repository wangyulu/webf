package middleware

import (
	"fmt"
	"geek/webf/framework"
	"net/http"
)

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre recovery")

		defer func() {
			if p := recover(); p != nil {
				c.Json(http.StatusInternalServerError, p)
			}
		}()

		err := c.Next()

		fmt.Println("middleware post recovery")

		return err
	}
}
