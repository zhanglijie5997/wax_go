package main

import (
	"go_study/redis"
	"go_study/router"
	"go_study/sql"
	os2 "go_study/utils/os"
)


func main() {
	os2.OsYaml()
	redis.ClientRedis()
	//web3.Web3Api()
	sql.Sql()
	router.Router()
}
