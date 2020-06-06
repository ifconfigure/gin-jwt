package Utils

import (
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

// JWT 签名结构
type JWTStruct struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "dgerfasdas1234234234"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID uint32 `json:"userId"`
	jwtgo.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWTStruct {
	return &JWTStruct{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

/**
 * 生成令牌
 */
func GenerateToken(userID uint32) (string, error) {
	j := JWTStruct{
		SigningKey: []byte("dgerfasdas1234234234"),
	}

	claims := CustomClaims{
		UserID: userID,

		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名生效时间
			ExpiresAt: time.Now().Unix() + 486400, // 过期时间(秒)
			Issuer:    "gin-jwt",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		fmt.Println(err)
		return token, err
	}

	return token, nil
}

// CreateToken 生成一个token
func (j *JWTStruct) CreateToken(claims CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWTStruct) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
