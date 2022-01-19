package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/model"
	"go_study/model/http_model"
	"go_study/sql"
	"net/http"
)

type Data struct {
	Content []model.User `json:"conent" `
}

func Index(c * gin.Context)  {
	var _user []model.User
	sql.DB.Table("users").Find(&_user)
	fmt.Println(_user)
	if _user != nil {
		c.JSON(http.StatusOK, http_model.Success{
			Code: http.StatusOK,
			Data: Data{
				Content: _user,
			},
			Message: "successage",
		})
	}else {
		c.JSON(http.StatusOK, http_model.HttpError{
			Code: http.StatusNotFound,
			Data: nil,
			Message: "error",
		})
	}
}