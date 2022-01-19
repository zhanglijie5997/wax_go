package middleware

import (
	"github.com/gin-gonic/gin"
	"go_study/config"
	"go_study/config/route_config"
	"go_study/model"
	"go_study/sql"
	"go_study/utils"
	"net/http"
	"time"
)
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// SetDBMiddleware 中间件验证token
func SetDBMiddleware(next *gin.Context)  {
	// get 请求不需要登录
	if next.Request.Method == http.MethodGet {
		next.Next()
		return
	}
	// 需要登录的接口
	_auth := []string{
		route_config.WaxRelation, route_config.WaxRelationUpdate,
		//route_config.HomeApi, route_config.LoginApi, route_config.Register,
	}
	_token := next.Request.Header.Get("token")
	_path := next.Request.URL.Path
	if _token != "" && IsContain(_auth, _path){
		var _userMsg model.UserMsg
	 	_sqlRes := sql.DB.Table("users").Where("token = ?", _token).First(&_userMsg)
		if _sqlRes.Error != nil {
			next.JSON(http.StatusOK, gin.H{
				"code": config.NeedLogin,
				"message": "Please login first",
			})
			// 未登录，直接中断请求
			next.Abort()
			//panic(_sqlRes.Error)
		}else {
			if _token == _userMsg.Token {
				// 判断token是否过期
				_t, err := utils.ParseToken(_token)
				if err != nil {
					next.JSON(http.StatusOK, gin.H{
						"code": config.TokenExpired,
						"message": "Token Resolution failure",
					})
					//panic(err)
				}else {
					if time.Now().Unix() > _t.ExpiresAt {
						next.JSON(http.StatusOK, gin.H{
							"code": config.TokenExpired,
							"message": "token is expired",
						})
						//panic("token is expired")
					}

					next.Next()
				}
			}else {
				next.JSON(http.StatusOK, gin.H{
					"code": config.NeedLogin,
					"message": "Log on first",
				})
				// 未登录，直接中断请求
				next.Abort()
			}

		}



	}else  {
		next.Next()
	}
	//url := next.Request.URL.String()
	//fmt.Println(router.HomeApi)
	//if url != "/" {
	//	next.JSON(http.StatusOK, gin.H{
	//		"code": config.NeedLogin,
	//		"message": "Log on first",
	//	})
	//	// 未登录，直接中断请求
	//	next.Abort()
	//}
	//gin.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
	//	ctx := context.WithValue(r.Context(), "DB", sql.DB.WithContext(timeoutContext))
	//	//next.ServeHTTP(w, r.WithContext(ctx))
	//	next.Next()
	//})

}