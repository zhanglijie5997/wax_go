package wax_list

import (
	"github.com/gin-gonic/gin"
	"go_study/config"
	"net/http"
)

func WaxList(c *gin.Context)  {
	email := c.Query("email")
	if email != "" {
		//res := utils.TryCatch(sql.DB.Table("users").Where())
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": config.EmailIsNotValidated,
			"message": "email is required",
		})
	}
}