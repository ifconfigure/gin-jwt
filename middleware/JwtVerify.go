package Middleware

import (
	Utils "gin-jwt/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(500 , map[string]interface{}{
				"code":500,
				"message":"no token",
			})
			c.Abort()
			return
		}

		j := Utils.NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == Utils.TokenExpired {
				c.JSON(500 , map[string]interface{}{
					"code":500,
					"message":"token expired",
				})
				c.Abort()
				return
			}
			c.JSON(500 , map[string]interface{}{
				"code":500,
				"message":err,
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("UserID", claims.UserID)
	}
}
