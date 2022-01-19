package http_model

type Success struct {
	Code int `json:"code"`
	Data interface{}  `json:"data"`
	Message string `json:"message"`
}

type HttpError struct {
	Code int 			`json:"code"`
	Data interface{} 	`json:"data"`
	Message string		`json:"message"`
}