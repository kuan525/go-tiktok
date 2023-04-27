package service

import (
	"common/dao"
	"common/log"
	"common/middleware"
	"common/models"
	"github.com/kataras/iris/v12"
	"sort"
	"time"
	"web_message/internal/request"
	"web_message/internal/response"
)

var Dao dao.Dao

// DouyinMessageChatHandler 聊天记录 当前登录用户和其他指定用户的聊天消息记录
func DouyinMessageChatHandler(ctx iris.Context, reqBody interface{}) {
	// 这里设定返回来回消息，并且后端排序好返回，前端直接按照顺序渲染即可
	req := reqBody.(*request.DouyinMessageChatReq)
	userId := middleware.GetUserIdByToken(req.Token)
	toUserId := req.ToUser

	userMessageList, err := Dao.GetMessageByFromUserIdAndPreMsgTime(userId, req.PreMsgTime)
	if err != nil {
		log.Logger.Infof(err.Error(), "从数据库获取数据失败")
		return
	}

	toUserMessageList, err := Dao.GetMessageByFromUserIdAndPreMsgTime(toUserId, req.PreMsgTime)
	if err != nil {
		log.Logger.Infof(err.Error(), "从数据库获取数据失败")
		return
	}

	var respMessageList []response.Message
	for _, modelMessage := range userMessageList {
		createTime := time.Unix(modelMessage.CreateTime, 0).Format("2006-01-02 15:04:05")
		respMessage := response.Message{
			Id:         modelMessage.Id,
			ToUserId:   modelMessage.ToUserId,
			FromUserId: modelMessage.FromUserId,
			Content:    modelMessage.Content,
			CreateTime: createTime,
		}
		respMessageList = append(respMessageList, respMessage)
	}

	for _, modelMessage := range toUserMessageList {
		createTime := time.Unix(modelMessage.CreateTime, 0).Format("2006-01-02 15:04:05")
		respMessage := response.Message{
			Id:         modelMessage.Id,
			ToUserId:   modelMessage.ToUserId,
			FromUserId: modelMessage.FromUserId,
			Content:    modelMessage.Content,
			CreateTime: createTime,
		}
		respMessageList = append(respMessageList, respMessage)
	}

	// 对消息进行排序,string比较默认按照rune比较，一个字符一个字符比较，字典顺，这里对字符比较标准，可以使用
	sort.Slice(respMessageList, func(i, j int) bool {
		return respMessageList[i].CreateTime < respMessageList[j].Content
	})

	err = ctx.JSON(response.DouyinMessageChatResp{
		StatusCode:  0,
		StatusMsg:   "获取成功",
		MessageList: respMessageList,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}

// DouyinMessageActionHandler 消息操作 登录用户对消息的相关操作，目前只支持消息发送
func DouyinMessageActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinMessageActionReq)
	if req.ActionType != 1 {
		log.Logger.Infof("前端相关功能尚未实现")
		err := ctx.JSON(response.DouyinMessageActionResp{
			StatusCode: 1,
			StatusMsg:  "功能尚未实现，尽情期待！",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "发送失败")
		}
		return
	}

	userId := middleware.GetUserIdByToken(req.Token)
	toUserId := req.ToUserId

	modelMessage := models.Message{
		FromUserId: userId,
		ToUserId:   toUserId,
		Content:    req.Content,
		CreateTime: time.Now().Unix(),
	}

	err := Dao.CreateMessage(&modelMessage)
	if err != nil {
		log.Logger.Infof(err.Error(), "插入失败")
		return
	}

	err = ctx.JSON(response.DouyinMessageActionResp{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}
