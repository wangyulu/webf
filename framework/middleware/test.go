package middleware

import (
	"fmt"
	"geek/webf/framework"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")

		err := c.Next()

		fmt.Println("middleware post test1")

		return err
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")

		err := c.Next()

		fmt.Println("middleware post test2")

		return err
	}
}

func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")

		err := c.Next()

		fmt.Println("middleware post test3")

		return err
	}
}
