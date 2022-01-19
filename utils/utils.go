package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Utils() string {
	fmt.Println("--.--")
	return "2123"
}

var JwtSecret = []byte("golang")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
/*
	CreateToken 生成token
	@param email 用户名
	@param password 密码
*/
func CreateToken(email string, password string) (string, error) {
	now := time.Now()
	end := now.Add(30 * time.Hour)
	claims := Claims{
		email,
		password,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: end.Unix(),
			// 指定token发行人
			Issuer:"gin-blog",
		},
	}
	_jwtScreet := []byte(JwtSecret)
	_tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=_tokenClaims.SignedString(_jwtScreet)
	return token,err
}
// ParseToken 解析jwt token
// @parma 	  t token
// @return
func ParseToken(t string) (*Claims, error) {
	_token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if _token != nil {
		if claims, ok := _token.Claims.(*Claims); ok && _token.Valid {
			return claims, nil
		}
	}
	return  nil, err
}

func TryCatch(db *gorm.DB, c *gin.Context) *gorm.DB {
	_res := db
	if _res.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": config.RequestError,
			"message": "request failed",
			"data": _res.Error,
		})
		return nil
	}else {
		return _res
	}
}