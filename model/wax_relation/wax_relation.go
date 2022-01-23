package wax_relation

import "go_study/model/wax_model"

type WaxRelationModel struct {
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	Total 	  int `json:"total"`
	Result	  []wax_model.WaxModel `json:"result"`
}