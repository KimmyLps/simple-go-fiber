package config

import (
	"fmt"

	m "test/models"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitialDb() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"test",
	)
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	DBConn.AutoMigrate(&m.Dogs{})
}
