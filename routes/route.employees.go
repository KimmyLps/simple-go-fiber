// +build ignore
package routes

import (
	c "github.com/KimmyLps/test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func RouteEmployee(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	})

	employee := v1.Group("/employees")
	employee.Get("/", c.GetEmployees)                 // GET http://localhost:3000/api/v1/employees
	employee.Get("/group", c.GetEmployeeGroupByAge)   // GET http://localhost:3000/api/v1/employees/group
	employee.Get("/:id", c.GetEmployee)               // GET http://localhost:3000/api/v1/employees/<employee_id>
	employee.Get("/search/:search", c.SearchEmployee) // GET http://localhost:3000/api/v1/employees/search/<search-value>
	employee.Post("/", auth, c.CreateEmployee)        // POST http://localhost:3000/api/v1/employees
	employee.Put("/:id", auth, c.UpdateEmployee)      // PUT http://localhost:3000/api/v1/employees/<employee_id>
	employee.Delete("/:id", auth, c.RemoveEmployee)   // DELETE http://localhost:3000/api/v1/employees/<employee_id>
}
