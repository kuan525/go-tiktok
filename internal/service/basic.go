package service

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
)

func DouyinFeedHandler(ctx iris.Context) error {
	var err error
	return err
}

func DouyinUserRegisterHandler(ctx iris.Context) error {
	token, err := middleware.GenerateToken(12)
	if err != nil {
		ctx.WriteString("注册失败")
	} else {
		ctx.WriteString(token)
	}

	return err
}

func DouyinUserLoginHandler(ctx iris.Context) error {
	var err error
	auth, _ := ctx.Values().Get("Auth").(*middleware.MyReq)
	err = ctx.JSON(iris.Map{
		"Token":  auth.Token,
		"UserId": auth.UserId,
	})

	return err
}

func DouyinUserHandler(ctx iris.Context) error {
	var err error
	return err
}

func DouyinPublishActionHandler(ctx iris.Context) error {
	var err error
	return err
}

func DouyinPublishListHandler(ctx iris.Context) error {
	var err error
	return err
}
