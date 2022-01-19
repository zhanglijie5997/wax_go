package model

type User struct {
	ID   		int		`json:"id"`
	Name 		string `json:"name"`
	Sex  		int    `json:"sex"`
	Email 		string `json:"email" form:"email" valid:"Required"`
	Password 	string `json:"password" form:"password" valid:"Required"`
	Uuid		string `json:"uuid" form:"uuid"`
	Token		string	`json:"token"`
}

type UserMsg struct {
	Name 		string `json:"name"`
	Id   		string `json:"id"`
	Sex  		int    `json:"sex"`
	Email 		string `json:"email" form:"email"`
	Uuid		string `json:"uuid" form:"uuid"`
	Token		string `json:"token" form:"token"`
}

