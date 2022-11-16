package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/WebServ/pkg"
	"time"
)

func SetupM1Routes(app *fiber.App) error {
	m1 := app.Group("/m1")

	m1.Post("/u1/currency", u1Route)
	m1.Post("/u2/calc", u2Route)
	m1.Post("/u3/calc", u3Route)
	m1.Post("/u4/calc", u4Route)

	return nil
}

func u1Route(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "multipart/form-data" && contentType != "application/x-www-form-urlencoded" {
		return pkg.BadRequest("Content-Type must be multipart/form-data")
	}

	data := new(struct {
		Currency float32 `form:"currency"`
	})

	if err := c.BodyParser(data); err != nil {
		return pkg.BadRequest(err.Error())
	}

	sek := data.Currency * 9.7

	return c.Render("m1/u1/currency", fiber.Map{
		"Dollar": data.Currency,
		"Sek":    sek,
	})

}

func u2Route(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "multipart/form-data" && contentType != "application/x-www-form-urlencoded" {
		return pkg.BadRequest("Content-Type must be multipart/form-data")
	}

	data := new(struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	})

	if err := c.BodyParser(data); err != nil {
		return pkg.BadRequest(err.Error())
	}

	years := 65 - data.Age

	return c.Render("m1/u2/calc", fiber.Map{
		"Name":  data.Name,
		"Years": years,
	})
}

func u3Route(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "multipart/form-data" && contentType != "application/x-www-form-urlencoded" {
		return pkg.BadRequest("Content-Type must be multipart/form-data")
	}

	data := new(struct {
		N1 int `form:"n1"`
		N2 int `form:"n2"`
	})

	if err := c.BodyParser(data); err != nil {
		return pkg.BadRequest(err.Error())
	}

	result := data.N1 + data.N2

	return c.Render("m1/u3/calc", fiber.Map{
		"Result": result,
		"N1":     data.N1,
		"N2":     data.N2,
	})
}

func u4Route(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "multipart/form-data" && contentType != "application/x-www-form-urlencoded" {
		return pkg.BadRequest("Content-Type must be multipart/form-data")
	}

	data := new(struct {
		Name      string `form:"name"`
		BirthYear int    `form:"birthyear"`
	})

	if err := c.BodyParser(data); err != nil {
		return pkg.BadRequest(err.Error())
	}

	t := time.Now()

	age := t.Year() - data.BirthYear

	return c.Render("m1/u4/calc", fiber.Map{
		"Name": data.Name,
		"Age":  age,
	})
}
