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

	if _, ok := GetUserIdAndValidByToken(tokenString); !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(&MyRespErr{
			StatusCode: 0,
			StatusMsg:  "Token不合法",
		})
		return
	}
}

// GetTokenString 获取GetTokenString
func GetTokenString(ctx iris.Context) string {
	tokenString := ""

	// 从HTTP请求头中获取JWT Token
	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		tokenString = authHeader[len("Bearer "):]
	}

	// 从HTTP请求体中获取JWT Token
	if tokenString == "" && ctx.Method() != "GET" {
		tokenString = ctx.FormValue("token")
	}

	return tokenString
}

// GetUserIdAndValidByToken 解析Token，得到用户id并判断是否合法
func GetUserIdAndValidByToken(tokenString string) (int64, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodES256.Alg() {
			err := fmt.Errorf("无效的签名算法：%v", token.Header["alg"])
			conf.HandleLogsErr(err, "")
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
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        uuid.NewString(),
		},
	})

	// 使用预共享密钥签名Token
	tokenString, err := token.SignedString([]byte("my_secret_key"))

	return tokenString, err
}
