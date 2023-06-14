package inertia

import (
	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) func(c *fiber.Ctx) error {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		c.Set("Vary", "X-Inertia")

		if c.Get("X-Inertia") == "" {
			return c.Next()
		}

		if c.Method() == "GET" && c.Get("X-Inertia-Version") != version() {
			c.Set("X-Inertia-Location", c.Path())
			c.Status(fiber.StatusConflict)
		}

		return c.Next()
	}
}
