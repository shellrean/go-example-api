package main

import (
	"github.com/gofiber/fiber/v2"
	"go-simple/api"
)

func main() {
	app := fiber.New()

	user := api.UserApi{}
	app.Get("/users", user.Index)

	_ = app.Listen(":9090")
}
