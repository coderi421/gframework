package middlewares

import "github.com/gin-gonic/gin"

// AuthStrategy 认证策略
type AuthStrategy interface {
	// AuthFunc 相当于是实现了一个名为 AuthFunc 的中间件
	AuthFunc() gin.HandlerFunc
}

// AuthOperator 认证实体工厂类
type AuthOperator struct {
	strategy AuthStrategy
}

func (ao *AuthOperator) SetStrategy(s AuthStrategy) {
	ao.strategy = s
}

// AuthFunc 返回一个 gin.HandlerFunc 中间件接口函数
func (ao *AuthOperator) AuthFunc() gin.HandlerFunc {
	return ao.strategy.AuthFunc()
}
