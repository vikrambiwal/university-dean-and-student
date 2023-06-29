package utility

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var database *DB

//  dsn := "root:root@tcp(localhost:3306)/udas?charset=utf8mb4&parseTime=True&loc=Local"
// database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// // database, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/udas?charset=utf8mb4&parseTime=True&loc=Local")
// // if err != nil {
// //     panic("Failed to connect to the database")
// // }

func Database() *gorm.DB {

	dsn := "root:root@tcp(localhost:3306)/udas?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("Failed to connect to the database")
	}

	return database
}
