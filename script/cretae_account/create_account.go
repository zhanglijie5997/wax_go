package cretae_account

import (
	"bytes"
	"fmt"
	"go_study/model/wax_model"
	os2 "go_study/utils/os"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)


func CreateAccount(db *gorm.DB) {
	_ymal := os2.YamlResult
	_path, err := os.Getwd()
	if err == nil {
		_path += _ymal.MySql.CreatePath
	}
	accountTet, err := os.ReadFile(_path)
	if err == nil {
		fmt.Println(time.Now().String())
		var res []wax_model.WaxModel
		line := strings.Split(string(accountTet), "\n")
		for i, _line := range line {
			_t := strings.Replace(_line, "\n", "", -1)
			_t = strings.Replace(_line, "\r", "", -1)
			txtRes := strings.Split(string(_t), "----")
			fmt.Println(i, txtRes)
			if len(txtRes) == 2 {
				account := wax_model.WaxModel{
					UserId: "",
					Email: "",
					Password: txtRes[1] ,
					UserAccount: txtRes[0],
					Inited: 0,
					Status: 1,
					LandId: "",
					Tools: "",
					CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				}
				res = append(res, account)
			}else {
				fmt.Println("File parsing failure")
			}
		}
		var buffer bytes.Buffer
		sqlStatements := "insert into wax (" +
			"user_id, " +
			"email, " +
			"password, " +
			"user_account, " +
			"inited, " +
			"status, " +
			"land_id, " +
			"tools, " +
			"created_at) " +
			"values"
		if _, err := buffer.WriteString(sqlStatements); err != nil {
			panic(err)
		}

		for i, e := range res {
			if i == len(res)-1 {
				buffer.WriteString(fmt.Sprintf("('%s','%s','%s', '%s', %d, %d, '%s','%s', '%s');", e.UserId, e.Email, e.Password, e.UserAccount,
					e.Inited, e.Status, e.LandId, e.Tools, e.CreatedAt))
			} else {
				buffer.WriteString(fmt.Sprintf("('%s','%s','%s', '%s', %d, %d, '%s', '%s', '%s'),", e.UserId, e.Email, e.Password, e.UserAccount,
					e.Inited, e.Status, e.LandId, e.Tools, e.CreatedAt))
			}
		}
		_r := db.Exec(buffer.String())
		if _r.Error != nil {
			fmt.Println(_r.Error)
		}else {
			fmt.Println("insert success")
		}
	}else {
		fmt.Println(err)
	}
}
