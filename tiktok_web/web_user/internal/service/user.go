package service

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
)

// DouyinUserRegisterHandler 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
func DouyinUserRegisterHandler(ctx iris.Context) error {
	token, err := middleware.GenerateToken(12)
	if err != nil {
		ctx.WriteString("注册失败")
	} else {
		ctx.WriteString(token)
	}

	return err
}

// DouyinUserLoginHandler 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
func DouyinUserLoginHandler(ctx iris.Context) error {
	var err error
	auth, _ := ctx.Values().Get("Auth").(*middleware.MyReq)
	err = ctx.JSON(iris.Map{
		"Token":  auth.Token,
		"UserId": auth.UserId,
	})

	return err
}

// DouyinUserHandler 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
func DouyinUserHandler(ctx iris.Context) error {
	var err error
	return err
}
