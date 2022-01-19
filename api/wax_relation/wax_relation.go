package wax_relation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/sql"
	"go_study/utils"
	"net/http"
)

type WaxRelationStuct struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Password string `json:"password"`
	UserAccount string	`json:"user_account"`
	Inited	int 	`json:"inited"`
	Status  int		`json:"status"`
	LandId	string 	`json:"land_id"`
	Tools	string 	`json:"tools"`
}
func WaxRelation(c *gin.Context)  {
	_email := c.Query("email")
	var list []WaxRelationStuct
	res := utils.TryCatch(sql.DB.Table("wax").Where("user_id = ?", _email).Find(&list), c)
	if res == nil {
		fmt.Println(res)
		panic("sql search error")
	}else {
		fmt.Println(list)
		c.JSON(http.StatusOK, gin.H{
			"code": config.Success,
			"message": "success",
			"data": list,
		})
		//panic("sql search success")
	}
}