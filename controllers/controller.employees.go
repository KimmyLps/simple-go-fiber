// +build ignore
package controllers

import (
	"github.com/KimmyLps/test/config"
	m "github.com/KimmyLps/test/models"

	"github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	db := config.DBConn
	var employees []m.Employee

	db.Find(&employees)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    employees,
	})
}

func GetEmployee(c *fiber.Ctx) error {
	db := config.DBConn
	id := c.Params("id")
	var employee m.Employee

	result := db.Find(&employee, "employee_id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "employee not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success",
		"data":    employee,
	})
}

func CreateEmployee(c *fiber.Ctx) error {
	db := config.DBConn
	var employee m.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	db.Create(&employee)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "success",
		"data":    employee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	db := config.DBConn
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "failed",
			"data":    nil,
		})
	}
	var employee m.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  503,
			"message": "failed",
			"data":    err,
		})
	}

	result := db.Where("employee_id = ?", id).Updates(employee)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "failed",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    nil,
	})
}

func RemoveEmployee(c *fiber.Ctx) error {
	db := config.DBConn
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "failed",
			"data":    nil,
		})
	}
	var employee m.Employee
	result := db.Where("employee_id = ?", id).Delete(&employee)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "employee not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    nil,
	})
}

func GetEmployeeGroupByAge(c *fiber.Ctx) error {
	db := config.DBConn
	var employees []m.Employee
	db.Find(&employees)

	groupedEmployees := groupByAge(employees)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    groupedEmployees,
	})
}

func groupByAge(employees []m.Employee) map[string][]m.Employee {
	ageGroup := make(map[string][]m.Employee)
	for _, employee := range employees {
		if employee.Age == nil {
			continue
		}

		age := *employee.Age
		if age < 24 {
			ageGroup["GenZ"] = append(ageGroup["GenZ"], employee)
		} else if age >= 24 && age <= 41 {
			ageGroup["GenY"] = append(ageGroup["GenY"], employee)
		} else if age >= 42 && age <= 56 {
			ageGroup["GenX"] = append(ageGroup["GenX"], employee)
		} else if age >= 67 && age <= 75 {
			ageGroup["Baby Boomer"] = append(ageGroup["Baby Boomer"], employee)
		} else if age > 75 {
			ageGroup["G.I. Generation"] = append(ageGroup["G.I. Generation"], employee)
		}
	}
	return ageGroup
}

func SearchEmployee(c *fiber.Ctx) error {
	db := config.DBConn
	search := c.Params("search")
	var employees []m.Employee

	results := db.Where("employee_id LIKE ? or name LIKE ? or last_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&employees)
	if results.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNoContent,
			"message": "no content",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    employees,
	})
}
