package config 

import (
	"gorm.io/gorm" 
	"gorm.io/driver/mysql"
)

var (
	db  gorm.DB 
)

func Connect(){
	dsn := "root:password123@tcp(127.0.0.1:3306)/moviesdb?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	db = *database
}

func GetDB() *gorm.DB{
	return &db 
}