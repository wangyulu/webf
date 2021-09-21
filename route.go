package main

import (
	"geek/webf/framework"
	"geek/webf/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Get("/foo", framework.TimeoutHandler(FooControllerHandler, 5))

	// HTTP Method + 静态路由
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缘
	subjectApi := core.Group("/subject")
	{
		// 为整个组添加中间件
		subjectApi.Use(middleware.Test3())

		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
