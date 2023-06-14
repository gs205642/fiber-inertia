package inertia

import "github.com/gofiber/fiber/v2"

// Config defines the config for middleware.
type Config struct {
	AssetUrl string
	RootView string
	Next     func(c *fiber.Ctx) bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	RootView: "app",
	Next:     nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.RootView == "" {
		cfg.RootView = ConfigDefault.RootView
	}

	return cfg
}
