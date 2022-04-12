package models


import (
	"gorm.io/gorm"
)

// 定义表结构
type (
	PersonModel struct {
		gorm.Model
		Name			string `json:"name"`
		Age				int `json:"age"`
		Gender			string `json:"gender"`
	}
)

// 定义model和数据库的映射关系
func (PersonModel) TableName() string {
	return "person"
}
