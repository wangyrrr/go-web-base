package entity

import "github.com/jinzhu/gorm"

type Demo struct {
	gorm.Model
	Name string
}
