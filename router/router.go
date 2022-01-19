package router

import (
	"github.com/gin-gonic/gin"
	"go_study/api/index"
	"go_study/api/login"
	register2 "go_study/api/register"
	"go_study/api/wax_list"
	"go_study/api/wax_relation"
	"go_study/api/wax_relation_update"
	"go_study/config"
	"go_study/config/route_config"
	"go_study/middleware"
)



func Router()  {
	r := gin.Default()
	//r.Static("/web", "../web/**/*")
	//r.LoadHTMLFiles("../web/**/*")
	r.Use(middleware.SetDBMiddleware)

	// 首页
	r.GET(route_config.HomeApi, index.Index)
	r.POST(route_config.HomeApi, config.ErrMethods)
	r.PUT(route_config.HomeApi, config.ErrMethods)
	r.DELETE(route_config.HomeApi, config.ErrMethods)

	// 登录
	r.POST(route_config.LoginApi, login.Login)
	r.GET(route_config.LoginApi, config.ErrMethods)
	r.PUT(route_config.LoginApi, config.ErrMethods)
	r.DELETE(route_config.LoginApi, config.ErrMethods)

	// 注册
	r.POST(route_config.Register, register2.Register)
	r.GET(route_config.Register, config.ErrMethods)
	r.PUT(route_config.Register, config.ErrMethods)
	r.DELETE(route_config.Register, config.ErrMethods)

	// 查询关联wax账号
	r.GET(route_config.WaxRelation, wax_relation.WaxRelation)
	r.POST(route_config.WaxRelation, config.ErrMethods)
	r.PUT(route_config.WaxRelation, config.ErrMethods)
	r.DELETE(route_config.WaxRelation, config.ErrMethods)

	// 更新关联wax账号
	r.GET(route_config.WaxRelationUpdate, config.ErrMethods)
	r.POST(route_config.WaxRelationUpdate, wax_relation_update.WaxRelationUpdate)
	r.PUT(route_config.WaxRelationUpdate, config.ErrMethods)
	r.DELETE(route_config.WaxRelationUpdate, config.ErrMethods)

	// 查询账号关联wax账号
	r.GET(route_config.WaxList, wax_list.WaxList)
	r.POST(route_config.WaxList, wax_list.WaxList)
	r.PUT(route_config.WaxList, wax_list.WaxList)
	r.DELETE(route_config.WaxList, wax_list.WaxList)

	r.Run(":6000")
}
