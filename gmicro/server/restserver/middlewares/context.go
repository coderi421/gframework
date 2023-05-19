package middlewares

import "github.com/gin-gonic/gin"

const (
	UsernameKey = "username"
	KeyUserID   = "userID"
	UserIP      = "ip"
)

// Context 请求上下文中间件，用于将一些常用的信息放入到gin.Context中
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从c中获取到ip地址
		//TODO 自己扩展
		//c.Set(UserIP, c.GetString(UsernameKey))
		c.Next()
	}
}
