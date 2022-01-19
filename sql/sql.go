package sql

import (
	"fmt"
	os2 "go_study/utils/os"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"time"
)



type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Sex  int    `json:"sex"`
}

type Order struct {
	OrderId int 		`json:"order_id"`
	Price   float32		`json:"price"`
}

var DB *gorm.DB
func Sql()  {
	_yaml := os2.YamlResult
	// 读取配置
	_path, err := os.Getwd()
	if err == nil {
		_path += "/pubspec.yaml"
	}
	yamlFile, err := ioutil.ReadFile(_path)
	if err == nil {
		fmt.Println()
		yaml.Unmarshal(yamlFile, &_yaml)
		fmt.Println(_yaml.MySql.Db)
	}else {
		fmt.Println(err, "错误")

	}
	// 读取配置
	sqlPath := _yaml.MySql.DbUser + ":" + _yaml.MySql.DbPassword + "@tcp(" + _yaml.MySql.DbHost + ":" + _yaml.MySql.DbPort +")/" + _yaml.MySql.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(sqlPath)
	db, _err := gorm.Open(mysql.Open(sqlPath), &gorm.Config{
		PrepareStmt: true, // 预编译
		Logger: logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true, // 跳过事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		NowFunc: func() time.Time {
			// 更改创建时间使用的函数
			return time.Now().Local()
		},
	})
	var user []User
	var order []Order
	if _err == nil {
		fmt.Println("db connecting success.")
		// 测试sql
		db.Table("users").Session(&gorm.Session{QueryFields: true}).Find(&user)
		//if user != nil {
		//	fmt.Println("user ---->>>> ",len(user) )
		//	for _,v := range user {
		//		fmt.Println("name -->>>",  v.Name)
		//		fmt.Println("sex -->>>",  v.Id)
		//		fmt.Println("id -->>>",  v.Sex)
		//	}
		//}else {
		//	fmt.Println("查询错误 --->>> ")
		//}
		fmt.Println(order, len(order))
		DB = db
	}
}