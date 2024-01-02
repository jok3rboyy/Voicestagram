package Repositories

import (
	"./types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseOpen() {
	dsn := "root:Harvey1994!!@tcp(localhost:3306)/ip_insta_opdracht?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&types.Post{}, &types.User{})

	Database = db

}
