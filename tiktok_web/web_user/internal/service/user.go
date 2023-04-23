package service

import (
	"common/conf"
	"common/log"
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_user/internal/dao"
	"web_user/internal/request"
	"web_user/internal/response"
)

var userDao dao.UserDao

// DouyinUserRegisterHandler 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
func DouyinUserRegisterHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinUserRegisterReq)

	if ok := userDao.UserIsExistByUsername(req.UserName); ok {
		ctx.StatusCode(iris.StatusOK)
		err := ctx.JSON(response.DouyinUserRegisterResp{
			StatusCode: 1,
			StatusMsg:  "当前用户已存在",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "service:发送resp失败")
		}
		return
	}

	// 生成雪花id，使用端口作为当前机器参数
	UserId := middleware.GetSnowflakeId(conf.Cfg.HttpAddr.Port)

	//生成token
	token, err := middleware.GenerateToken(UserId)
	if err != nil {
		log.Logger.Infof(err.Error(), "service:创建token错误")
	}

	if ok := userDao.Register(req.UserName, req.Password, UserId); !ok {
		ctx.StatusCode(iris.StatusInternalServerError)
		err := ctx.JSON(response.DouyinUserRegisterResp{
			StatusCode: 1,
			StatusMsg:  "注册失败",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "注册失败")
		}
		return
	}

	ctx.StatusCode(iris.StatusOK)
	err = ctx.JSON(response.DouyinUserRegisterResp{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     UserId,
		Token:      token,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "service:发送resp失败")
	}
}

// DouyinUserLoginHandler 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
func DouyinUserLoginHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinUserLoginReq)

	if ok := userDao.UserIsExistByUsername(req.UserName); !ok {
		ctx.StatusCode(iris.StatusBadRequest)
		err := ctx.JSON(response.DouyinUserLoginResp{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "service:发送resp失败")
		}
		return
	}

	token, userId, ok := userDao.GetTokenAndUserIdByUsernameAndPassword(req.UserName, req.Password)
	if !ok {
		ctx.StatusCode(iris.StatusBadRequest)
		err := ctx.JSON(response.DouyinUserLoginResp{
			StatusCode: 1,
			StatusMsg:  "密码错误",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "service:发送resp失败")
		}
		return
	}

	ctx.StatusCode(iris.StatusOK)
	err := ctx.JSON(response.DouyinUserLoginResp{
		StatusCode: 0,
		StatusMsg:  "登陆成功",
		Token:      token,
		UserId:     userId,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "service:发送resp失败")
	}
}

// DouyinUserHandler 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
func DouyinUserHandler(ctx iris.Context, reqBody interface{}) {
	userId, err := ctx.URLParamInt64("user_id")

	user, ok := userDao.GetUserByUserId(userId)
	if !ok || err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		err := ctx.JSON(response.DouyinUserResp{
			StatusCode: 1,
			StatusMsg:  "获取信息失败",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "service:发送resp失败")
		}
		return
	}

	ctx.StatusCode(iris.StatusOK)
	err = ctx.JSON(response.DouyinUserResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		User:       *user,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "service:发送resp失败")
	}
}
