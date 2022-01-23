package wax_model

type WaxModel struct {
	UserId 		string  `json:"user_id" gorm:"column:user_id"`
	Email  		string	`json:"email" gorm:"column:email"`
	Password	string 	`json:"password" gorm:"column:password"`
	UserAccount string	`json:"user_account" gorm:"column:user_account"`
	Inited		int		`json:"inited" gorm:"column:inited"`
	Status		int		`json:"status" gorm:"column:status"`
	LandId		string	`json:"land_id" gorm:"column:land_id"`
	Tools		string	`json:"tools" gorm:"column:tools"`
	CreatedAt   string 	`json:"created_at" gorm:"column:created_at"`
	UpdatedAt 	string  `json:"updated_at" gorm:"column:updated_at"`
}