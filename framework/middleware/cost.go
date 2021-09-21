package middleware

import (
	"fmt"
	"geek/webf/framework"
	"time"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre cost")

		start := time.Now()

		err := c.Next()

		end := time.Now()

		fmt.Printf("middleware ing cost, api:%s, cost:%s \n", c.GetRequest().RequestURI, end.Sub(start))

		fmt.Println("middleware post cost")

		return err
	}
}
