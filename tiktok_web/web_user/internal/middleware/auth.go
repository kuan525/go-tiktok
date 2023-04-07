package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	//"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"go-tiktok/conf"
	"time"
)

type MyReq struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

type MyRespErr struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type MyClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

// Auth 鉴权
func Auth(ctx iris.Context) {
	tokenString := GetTokenString(ctx)
	if tokenString == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(&MyRespErr{
			StatusCode: 0,
			StatusMsg:  "未提供有效的Token",
		})
		return
	}

	userId, ok := GetUserIdAndValidByToken(tokenString)
	if !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(&MyRespErr{
			StatusCode: 0,
			StatusMsg:  "Token不合法",
		})
		return
	}
	ctx.Values().Set("Auth", &MyReq{
		Token:  tokenString,
		UserId: userId,
	})
	ctx.Next()
}

// GetTokenString 获取GetTokenString
func GetTokenString(ctx iris.Context) string {
	tokenString := ""
	// 从HTTP请求头中获取JWT Token
	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		tokenString = authHeader[len("Bearer "):]
	}

	var my MyReq
	if tokenString == "" {
		err := ctx.ReadJSON(&my)
		if err != nil {
			conf.HandleLogsErr(err, "body中获取token失败")
		}
		tokenString = my.Token
	}
	return tokenString
}

// GetUserIdAndValidByToken 解析Token，得到用户id并判断是否合法
func GetUserIdAndValidByToken(tokenString string) (int64, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			err := fmt.Errorf("无效的签名算法：%v", token.Header["alg"])
			conf.HandleLogsErr(err, "")
			return nil, err
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return 0, false
	}
	if !token.Valid {
		return 0, false
	}

	claims := token.Claims.(*MyClaims)
	return claims.UserId, true
}

// GenerateToken 生成token
func GenerateToken(userId int64) (string, error) {
	// 创建JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "kuan525",
			Subject:   "tiktok用户",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(6 * time.Hour).Unix(),
			Id:        uuid.NewString(),
		},
	})

	// 使用预共享密钥签名Token
	tokenString, err := token.SignedString([]byte("my_secret_key"))

	return tokenString, err
}
