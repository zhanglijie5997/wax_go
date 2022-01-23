package wax_list

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/model/http_model"
	"go_study/model/wax_list_model"
	"go_study/model/wax_model"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"strconv"
)

func WaxList(c *gin.Context)  {
	usedId := c.Query("user_id")
	inited, _ := strconv.Atoi(c.Query("inited"))
	var list []wax_model.WaxModel
	var total int64
	if usedId != "" {
		//res := utils.TryCatch(sql.DB.Table("users").Where())
		_db := utils.TryCatch(
				sql.DB.Table("wax").Where("user_id = ?", usedId).Find(&list), c,
			)

		if _db != nil {
			_db.Count(&total)
			if c.Query("inited") != "" {
				_db.Where("inited = ?", inited).Find(&list)
				_db.Count(&total)
			}
		}
		c.JSON(http.StatusOK, http_model.HttpModel{
			Code: config.Success,
			Message: "success",
			Data: wax_list_model.WaxListModel{
				Data: list,
				Total: int(total),
			},
		})
		fmt.Println(list)
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": config.EmailIsNotValidated,
			"message": "email is required",
		})
	}
}