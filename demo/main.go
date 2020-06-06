package main

import (
	Middleware "gin-jwt/middleware"
	Utils "gin-jwt/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID uint32 `json:"user_id"`
}

func main() {

	user := User{ID: 10086}

	tokenString, _ := Utils.GenerateToken(user.ID)

	r := gin.Default()

	//生成token
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"user":  user,
			"token": tokenString,
		})
	})

	//校验并获取token
	r.GET("/user" ,Middleware.JWTAuth() , func(c *gin.Context) {

		userID := c.MustGet("UserID").(uint32)

		c.JSON(200, map[string]interface{}{
			"user_id":  userID,
		})
	})

	_ = r.Run(":8081")
}
