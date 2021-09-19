package main

import (
	"geek/webf/framework"
	"net/http"
)

func SubjectListController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectListController")
	return nil
}

func SubjectAddController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectAddController")
	return nil
}

func SubjectDelController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectGetController")
	return nil
}

func SubjectNameController(ctx *framework.Context) error {
	ctx.Json(http.StatusOK, "ok, SubjectNameController")
	return nil
}
