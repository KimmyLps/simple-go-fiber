// +build ignore
package routes

import (
	"github.com/KimmyLps/test/controllers"
	c "github.com/KimmyLps/test/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteTest(app *fiber.App) {

	// 5.0
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"gofiber": "21022566",
	// 	},
	// }))

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Get("/inet", c.QueryTest)
	v1.Post("/register", c.Register)

	v2 := api.Group("/v2")
	v2.Get("/valid", c.ValidateTest)

	v3 := api.Group("/v3")
	v3.Get("/fac/:factorial", c.CalculateFac)    // 5.1
	v3.Get("/:tax_id", controllers.ConvertAscii) // 5.2
}
