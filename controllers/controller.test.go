package controllers

import (
	// "fmt"
	"log"
	// "regexp"
	"strconv"
	m "test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidateTest(c *fiber.Ctx) error {
	//Connect to database
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func CalculateFac(c *fiber.Ctx) error {
	// calculate the factorial of the number.
	p := c.Params("factorial")
	f, err := strconv.ParseInt(p, 10, 64)
	if err!= nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

	if f < 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "factorial cannot be negative",
        })
    }
	if f == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "factorial cannot be 0",
        })
    }
	return c.JSON(fiber.Map{
		"factorial": fac(f),
	})
}

func fac(n int64) int64  {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}


func ConvertAscii(c *fiber.Ctx) error  {
	taxId := c.Params("tax_id")

	ascii:=make([]int, len(taxId))
	for i := 0; i < len(taxId); i++ {
        ascii[i] = int(taxId[i])
    }

	return c.JSON(fiber.Map{
		"ascii": ascii,
	})
}

func Register(c *fiber.Ctx) error {
	//Connect to database
	company := new(m.Company)
	if err := c.BodyParser(&company); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(company)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.Error(),
		})
	}

	// regex := regexp.MustCompile(`^[a-z0-9-]{2,30}$`)
	return c.JSON(company)
}