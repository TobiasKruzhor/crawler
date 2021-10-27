package model

// 文章类型模型

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
