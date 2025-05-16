package midelware

import (
	"strings"
	"task-service/internal/dto"

	"github.com/gofiber/fiber/v2"
)

// key=Auth Value=admin;admin Add to Headers
func Auntification(users map[string]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		d := c.GetReqHeaders()

		if v, ok := d["Auth"]; ok {
			param := strings.Split(v[0], ";")
			user, token := param[0], param[1]
			if auth, ok := users[user]; ok {
				if auth == token {
					c.Next()
					return nil
				}
			}

		}
		return dto.BadResponseError(c, (fiber.StatusForbidden), "Доступ запрещен")
	}

}
