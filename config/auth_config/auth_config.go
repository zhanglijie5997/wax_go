package auth_config

import "go_study/config/route_config"



func Auth() []string {
	return []string{
		route_config.WaxRelation, route_config.WaxRelationUpdate,
		//route_config.HomeApi, route_config.LoginApi, route_config.Register,
	}
}