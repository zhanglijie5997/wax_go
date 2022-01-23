package wax_list_model

import "go_study/model/wax_model"

type WaxListModel struct {
	Data []wax_model.WaxModel `json:"data"`
	Total int `json:"total"`
}