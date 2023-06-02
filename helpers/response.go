package helpers

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

type Responses struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func EncryptPassword(text string) string {
	passwordHash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(passwordHash[:])
}

func CustomResponse(ctx *fiber.Ctx, data interface{}, message interface{}, statusCode int) error {

	return ctx.Status(statusCode).JSON(&Responses{
		Data:    data,
		Message: message,
	})
}
