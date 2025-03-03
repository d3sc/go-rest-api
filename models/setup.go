package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// setup ke database
func ConnectDatabase() {
	// koneksi ke database
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_rest_api"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
