package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/model"
	"go_study/sql"
	"go_study/utils"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
)


type LoginData struct {
	Email     string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context)  {
	emailRange := regexp.MustCompile(`(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@(([[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	email := c.PostForm("email")
	password := c.PostForm("password")
	emailResult := len(emailRange.FindAllString(email, -1))
	var user model.User
	if emailResult == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": config.EmailIsNotValidated,
			"message": "email is not validated",
		})
	}else  if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": config.PasswordIsNotValidate,
			"message": "password is not validated",
		})
	}else {
		// 根据email 查询用户信息
		sql.DB.Table("users").Session(&gorm.Session{QueryFields: true}).Where("email", email).First(&user)
		fmt.Println(user)
		if user.Email != "" {
			if user.Password == password {
				var user model.User
				_res := sql.DB.Table("users").First(&user, "email = ?", email)
				if _res.Error != nil {
					c.JSON(http.StatusOK, gin.H{
						"code": config.LoginFailed,
						"message": "login failed",
						"data": _res.Error,
					})
					panic(_res.Error)
				}else  {
					// 登录，生成token
					_token, err := utils.CreateToken(email, password)
					if err != nil {
						panic(err)
					}
					_result := utils.TryCatch(sql.DB.Table("users").Where("email = ?", email).Update("token", _token), c)
					if _result == nil{
						panic(_result.Error)
					}else {
						c.JSON(http.StatusOK, gin.H{
							"code": config.Success,
							"message": "login success",
							"data": model.UserMsg{
								Email: email,
								Sex:   1,
								Id:    strconv.Itoa(user.ID),
								Name:  email,
								Uuid:  user.Uuid,
								Token: _token,
							},
						})
					}
				}
			}else {
				c.JSON(http.StatusOK, gin.H{
					"code": config.PasswordIsNotValidate,
					"message": "password is not validate",
				})
			}
		}else {
			c.JSON(http.StatusOK, gin.H{
				"code": config.IsNotRegister,
				"message": "email is not register",
			})
		}
	}
}
