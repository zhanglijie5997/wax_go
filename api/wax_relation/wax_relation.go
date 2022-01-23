package wax_relation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/model/http_model"
	"go_study/model/wax_model"
	"go_study/model/wax_relation"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"strconv"
)




func WaxRelation(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	//account := c.Query("account")
	var list []wax_model.WaxModel
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 8
	}
	if page > 0 && pageSize > 0 {
		r := utils.TryCatch(sql.DB.Table("wax").Limit(pageSize).Offset((page - 1) * pageSize).Find(&list), c)
		var total int64
		sql.DB.Table("wax").Count(&total)
		if r != nil {
			fmt.Println(list)
			c.JSON(http.StatusOK, http_model.HttpModel{
				Code: config.Success,
				Message: "success",
				Data: wax_relation.WaxRelationModel{
					PageNum: page,
					PageSize: pageSize,
					Total: int(total) ,
					Result: list,
				},
			})
		}
	}

}