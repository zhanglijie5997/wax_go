package wax_relation_update

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/middleware"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"strconv"
)

type WaxAccount struct {
	UserId string `json:"user_id"`
	Email  string  `json:"email"`
	Password string `json:"password"`
	UserAccount string `json:"user_account"`
	Inited int `json:"inited"`
	Status int `json:"status"`
	LandId string `json:"land_id"`
	Tools string `json:"tools"`
}

func WaxRelationUpdate(c *gin.Context)  {
	inited := c.PostForm("inited")
	status := c.PostForm("status")
	landId := c.PostForm("land_id")
	tools := c.PostForm("tools")
	userId := c.PostForm("user_id")
	fmt.Println( inited, status, landId, tools, userId)

	if userId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": config.EmailIsNotValidated,
			"message": "user_id is not validate",
		})
		return
	}

	statusList := []string{"0", "1"}
	if status != "" && !middleware.IsContain(statusList, status) {
		 c.JSON(http.StatusOK, gin.H{
			"code": config.RequestError,
			"message": "status is must 0 or 1",
		 })
	}

	if inited != "" && !middleware.IsContain(statusList, inited) {
		c.JSON(http.StatusOK, gin.H{
			"code": config.RequestError,
			"message": "inited is must 0 or 1",
		})
	}
	var userAccount WaxAccount
	_res := utils.TryCatch(sql.DB.Table("wax").Where("user_id = ?", userId).First(&userAccount), c)
	if _res != nil {
		fmt.Println(userAccount)
		if landId != "" {
			userAccount.LandId = landId
		}
		if _inited, err := strconv.Atoi(inited); err == nil {
			userAccount.Inited = _inited
		}
		if _status, err := strconv.Atoi(status); err == nil {
			userAccount.Status = _status
		}
		if tools != "" {
			userAccount.Tools = tools
		}
		// 更新账号数据
		_result := utils.TryCatch(sql.DB.Table("wax").Where("user_id = ?", userId).Updates(&userAccount), c)
		if _result != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": config.Success,
				"message": "scuccess",
				"data": userAccount,
			})
		}

	}
	//c.JSON(http.StatusOK, gin.H{
	//	"code": config.Success,
	//	"message": "success",
	//})
}
