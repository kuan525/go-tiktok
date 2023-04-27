package service

import (
	"common/dao"
	"common/log"
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_relation/internal/request"
	"web_relation/internal/response"
)

var Dao dao.Dao

// DouyinRelationActionHandler 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
func DouyinRelationActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinRelationActionReq)
	userId := middleware.GetUserIdByToken(req.Token)

	if req.ActionType == 1 { // 关注
		err := Dao.AddRelation(userId, req.ToUser)
		if err != nil {
			log.Logger.Infof(err.Error(), "数据库中关注失败")
			return
		}
	} else { // 取消
		err := Dao.DeleteRelation(userId, req.ToUser)
		if err != nil {
			log.Logger.Infof(err.Error(), "数据库中取消关注失败")
			return
		}
	}

	err := ctx.JSON(response.DouyinRelationActionResp{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}

// DouyinRelationFollowListHandler 用户关注列表 登录用户关注的所有用户列表。
func DouyinRelationFollowListHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinRelationFollowListReq)
	userId := req.UserId
	followList, err := Dao.GetFollowListByUserId(req.UserId)
	if err != nil {
		log.Logger.Infof(err.Error(), "数据库获取数据失败")
		return
	}

	var respFollowUserList []response.User
	for _, followUserId := range followList {
		modelUser, ok := Dao.GetUserByUserId(int64(followUserId))
		if !ok {
			log.Logger.Infof(err.Error(), "通过userId获取user失败")
			return
		}
		respUser := response.User{
			Id:              modelUser.UserId,
			Name:            modelUser.Name,
			FollowCount:     modelUser.FollowCount,
			FollowerCount:   modelUser.FollowerCount,
			Avatar:          modelUser.Avatar,
			BackgroundImage: modelUser.BackgroundImage,
			Signature:       modelUser.Signature,
			WorkCount:       modelUser.WorkCount,
			// 这里是当前访问的用户是否对改评论的用户关注，不是对视频作者
			IsFollow:       Dao.IsFollow(modelUser.UserId, userId),
			TotalFavorited: Dao.GetNumUserAllGetFavorite(modelUser.UserId),
			FavoriteCount:  Dao.GetNumFavorite(modelUser.UserId),
		}
		respFollowUserList = append(respFollowUserList, respUser)
	}

	err = ctx.JSON(response.DouyinRelationFollowListResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		UserList:   respFollowUserList,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}

// DouyinRelationFollowerListHandler 用户粉丝列表 所有关注登录用户的粉丝列表。
func DouyinRelationFollowerListHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinRelationFollowerListReq)
	userId := req.UserId
	followerList, err := Dao.GetFollowerListByUserId(req.UserId)
	if err != nil {
		log.Logger.Infof(err.Error(), "数据库获取数据失败")
		return
	}

	var respFollowerUserList []response.User
	for _, followUserId := range followerList {
		modelUser, ok := Dao.GetUserByUserId(int64(followUserId))
		if !ok {
			log.Logger.Infof(err.Error(), "通过userId获取user失败")
			return
		}
		respUser := response.User{
			Id:              modelUser.UserId,
			Name:            modelUser.Name,
			FollowCount:     modelUser.FollowCount,
			FollowerCount:   modelUser.FollowerCount,
			Avatar:          modelUser.Avatar,
			BackgroundImage: modelUser.BackgroundImage,
			Signature:       modelUser.Signature,
			WorkCount:       modelUser.WorkCount,
			// 这里是当前访问的用户是否对改评论的用户关注，不是对视频作者
			IsFollow:       Dao.IsFollow(modelUser.UserId, userId),
			TotalFavorited: Dao.GetNumUserAllGetFavorite(modelUser.UserId),
			FavoriteCount:  Dao.GetNumFavorite(modelUser.UserId),
		}
		respFollowerUserList = append(respFollowerUserList, respUser)
	}

	err = ctx.JSON(response.DouyinRelationFollowListResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		UserList:   respFollowerUserList,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}

// DouyinRelationFriendListHandler 用户好友列表 所有关注登录用户的粉丝列表。
//func DouyinRelationFriendListHandler(ctx iris.Context, reqBody interface{}) {
//
//}
