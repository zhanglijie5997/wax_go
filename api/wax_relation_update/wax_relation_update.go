package wax_relation_update

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/middleware"
	"go_study/model/http_model"
	"go_study/model/wax_model"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"strconv"
	"time"
)



func WaxRelationUpdate(c *gin.Context)  {
	inited := c.PostForm("inited")
	status := c.PostForm("status")
	landId := c.PostForm("land_id")
	tools := c.PostForm("tools")
	userId := c.PostForm("user_id")
	account := c.PostForm("user_account")
	fmt.Println( inited, status, landId, tools, userId)
	// 4km10c@vip601.cn
	if account == "" {
		c.JSON(http.StatusOK, http_model.HttpModel{
			Code: config.EmailIsNotValidated,
			Message: "account is required",
		})
		return
	}

	//if userId != "" {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": config.EmailIsNotValidated,
	//		"message": "user_id is not validate",
	//	})
	//	return
	//}

	statusList := []string{"0", "1"}
	if status != "" && !middleware.IsContain(statusList, status) {
		 c.JSON(http.StatusOK, http_model.HttpModel{
			 Code: config.RequestError,
			 Message: "status is must 0 or 1",
		 })
	}

	if inited != "" && !middleware.IsContain(statusList, inited) {
		c.JSON(http.StatusOK, http_model.HttpModel{
			Code:  config.RequestError,
			Message: "inited is must 0 or 1",
		} )
	}
	var userAccount wax_model.WaxModel
	_res := utils.TryCatch(sql.DB.Table("wax").Where("user_account = ?", account).First(&userAccount), c)
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
		if userId != "" {
			userAccount.UserId = userId;
		}
		userAccount.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		// 更新账号数据
		_result := utils.TryCatch(sql.DB.Table("wax").Where("user_account = ?", account).Updates(&userAccount), c)
		if _result != nil {
			c.JSON(http.StatusOK,
				http_model.HttpModel{
					Code: config.Success,
					Data: userAccount,
					Message: "scuccess",
				},
				//gin.H{
				//	"code": config.Success,
				//	"message": "scuccess",
				//	"data": userAccount,
				//}
			)
		}
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"code": config.Success,
	//	"message": "success",
	//})
}
