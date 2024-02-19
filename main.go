package main

import (
	"test/config"
	r "test/routes"

	"github.com/gofiber/fiber/v2"
)

// func initialDb() {
// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
// 		"root",
// 		"",
// 		"127.0.0.1",
// 		"3306",
// 		"test",
// 	)
// 	var err error
// 	config.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Database connected!")

// }
func main() {
	app := fiber.New()

	config.InitialDb()

	r.RouteTest(app)

	app.Listen(":3000")

	// test
}
