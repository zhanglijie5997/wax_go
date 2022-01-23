package wax_relation_create

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/model/http_model"
	"go_study/model/wax_model"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"strconv"
	"time"
)

func WaxRelationCreate(c *gin.Context)  {
	var params = []string{"user_id", "email", "password", "user_account", "inited", "status", "land_id", "tools"}
	var waxAccount wax_model.WaxModel
	for i, s := range params {
		fmt.Println(i, s, c.PostForm(s))
		r := c.PostForm(s)
		switch s {
			case "user_id": waxAccount.UserId = r
			case "email": waxAccount.Email = r
			case "password": waxAccount.Password = r
			case "user_account":waxAccount.UserAccount = r
			case "inited":
				res, _err := strconv.Atoi(r)
				if _err != nil {
					waxAccount.Inited = 0
				}else {
					waxAccount.Inited = res
				}

			case "status":
				res, _err := strconv.Atoi(r)
				if _err != nil {
					waxAccount.Status = 0
				}else {
					waxAccount.Status = res
				}
			case "land_id": waxAccount.LandId = r
			case "tools": waxAccount.Tools = r
		}

		if waxAccount.UserAccount == "" {
			c.JSON(http.StatusOK, http_model.HttpModel{
				Code: config.RequestError,
				Message: "user_account is required",
			})
			return
		}

		if waxAccount.Password == "" {
			c.JSON(http.StatusOK, http_model.HttpModel{
				Code: config.RequestError,
				Message: "password is required",
			})
			return
		}

		waxAccount.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		res := utils.TryCatch(sql.DB.Table("wax").Create(&waxAccount), c)

		if res == nil {
			c.JSON(http.StatusOK, http_model.HttpModel{
				Code: config.Success,
				Message: "success",
			})
		}
	}
}