package middleware

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

func GetSnowflakeId(port string) int64 {
	// 将port的格式转换成int
	portInt, _ := strconv.Atoi(port)

	// 将端口号作为节点编号，来作为雪花id生成的参数
	node, _ := snowflake.NewNode(int64(portInt) - 6000)

	// 生成雪花id
	return node.Generate().Int64()
}
