# gin-jwt
Golang - Gin JWT Demo


### 1、Usage | 使用方式
```
go get github.com/ifconfigure/gin-jwt
```

### 2、Generate Token | 颁发Token
```
tokenString, _ := Utils.GenerateToken(user.ID)
```

### 3、Verify Token | 鉴权
Gin
```
r.GET("/user" ,Middleware.JWTAuth() , Controller)
```

