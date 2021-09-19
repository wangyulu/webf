package main

import (
	"geek/webf/framework"
)

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
