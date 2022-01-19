package register

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_study/config"
	"go_study/model"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"regexp"
	"strconv"
)

func Register(c *gin.Context)  {
	var userFind model.User
	var user model.User
	emailRange := regexp.MustCompile(`(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@(([[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" {
		 c.JSON(http.StatusOK, gin.H{
			"code": config.EmailIsNotValidated,
			"message": "email is not validated",
		})
	}else if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": config.PasswordIsNotValidate,
			"message": "password is not validated",
		})
	}else if email != "" && password != "" {
		if emailRange != nil {
			_emailList := emailRange.FindAllString(email, -1)
			_len := len(_emailList)
			if _len == 0  {
				c.JSON(http.StatusOK, gin.H{
					"code": config.EmailIsNotValidated,
					"message": "email is not validated",
				})
			}else {
				_uuid := uuid.New().String()
				// 查询是否存在用户
				sql.DB.Table("users").First(&userFind, "email = ?", email)
				if userFind.Email == "" {
					_token, err := utils.CreateToken(email, password)
					if err != nil {
						panic(err)
					}

					user = model.User{
						Email: email,
						Password: password,
						Sex: 1,
						Name: email,
						Uuid: _uuid,
						Token: _token,
					}
				    res := sql.DB.Table("users").Create(&user)

					if res.Error == nil {
						c.JSON(http.StatusOK, gin.H{
							"code": config.Success,
							"message": "register success",
							"data": model.UserMsg{
								Email: email,
								Sex:   1,
								Id: strconv.Itoa(user.ID),
								Name:  email,
								Uuid:  _uuid,
								Token: _token,
							},
						})
					}else {
						c.JSON(http.StatusOK, gin.H{
							"code": config.RegisterFailed,
							"message": "register failed",
							"data": res.Error,
						})
					}

				}else  {
					fmt.Println(userFind)
					c.JSON(http.StatusOK, gin.H{
						"code": config.IsRegistered,
						"message": "email is registed",
					})
				}

			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": config.EmailIsRegister,
				"message": "email is register",
			})
		}
	}
}