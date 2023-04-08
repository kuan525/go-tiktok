package service

import "github.com/kataras/iris/v12"

// DouyinFeedHandler 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
func DouyinFeedHandler(ctx iris.Context) error {
	var err error
	return err
}
