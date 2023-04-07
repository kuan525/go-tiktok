# go-tiktok
极简版抖音项目（字节跳动青训营）

# APIs
基础接口：

| Method | Path | Description               |
| :----- | ---- |---------------------------|
| GET    | /douyin/feed | 返回视频列表，可不登陆               |
| POST   | /douyin/user/register | 注册接口，提供用户名和密码             |
| POST   | /douyin/user/login | 登录接口，提供用户名和密码，返回用户id和token |
| GET   | /douyin/user | 获取登录用户的 id、昵称，关注数和粉丝数     |
| POST   | /douyin/publish/action | 视频投稿                      |
| GET   | /douyin/publish/list | 登录用户发布的视频列表               |

互动接口：

| Method | Path                 | Description |
| :----- |----------------------| ----------- |
| POST | /douyin/favorite/action | 登录用户对视频的点赞和取消点赞操作。 |
| GET | /douyin/favorite/list | 登录用户的所有点赞视频 |
| POST | /douyin/comment/action | 评论操作 |
| GET | /douyin/comment/list | 查看视频的所有评论，按发布时间排序 |

社交接口：

| Method | Path                        | Description            |
| :----- |-----------------------------|------------------------|
| POST | /douyin/relation/action     | 登陆用户或其他用户进行关注或取消关注     |
| GET | /douyin/relation/follow/list | 登陆用户关注的所有用户列表          |
| GET | /douyin/relation/follower/list | 所有关注登陆用户的粉丝列表          |
| GET | /douyin/relation/friend/list | 所有关注登录用户的粉丝列表及聊天消息。    |
| GET | /douyin/message/chat        | 当前登录用户和其他指定用户的聊天消息记录   |
| POST | /douyin/message/action      | 登录用户对消息的相关操作，目前只支持消息发送 |


# 架构图（初）
![架构图](https://s3.bmp.ovh/imgs/2023/04/08/6a933c18484b4744.jpg)


