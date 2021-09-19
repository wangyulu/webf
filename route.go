package main

import (
	"geek/webf/framework"
)

func registerRouter(core *framework.Core) {
	// core.Get("foo", FooControllerHandler)

	// HTTP Method + 静态路由
	core.Get("/user/login", UserLoginController)

	// 批量通用前缘
	subjectApi := core.Group("/subject")
	{
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
