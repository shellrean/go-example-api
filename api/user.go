package api

import (
	"github.com/gofiber/fiber/v2"
	"go-simple/dto"
)

type UserApi struct {
}

func (U UserApi) Index(ctx *fiber.Ctx) error {
	users := []dto.User{{Id: "1", Name: "Shellrean"}, {Id: "2", Name: "Senandika Selesa"}}
	return ctx.JSON(users)
}
