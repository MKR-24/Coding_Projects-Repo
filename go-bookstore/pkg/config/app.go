package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)
var(
	db *gorm.DB
)

func Connect(){
	d,err:=gorm.Open("mysql", "MKR:FoodisLove/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}