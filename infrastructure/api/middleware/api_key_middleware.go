package middleware

import (
	"strings"

	"github.com/deecodeid/api_nowted/config"
	"github.com/deecodeid/api_nowted/helper"
	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("x-api-key")
	if strings.Compare(apiKey, config.ENV.APIKey) != 0 {
		return helper.HandleResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}
	return c.Next()
}
