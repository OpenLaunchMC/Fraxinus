package authlibali

import (
	"github.com/gofiber/fiber/v2"

	ferrors "github.com/ProjectOpenLaunch/Fraxinus/errors"
)

type Config struct {
	// Next defines a function to skip this middleware.
	Next func(*fiber.Ctx) bool

	Location string

	ALIError ferrors.YggdrasilErrorResponse
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next: nil,
	ALIError: ferrors.YggdrasilErrorResponse{
		StatusCode: 400,
		Error: "IllegalArgumentException",
	},
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.ALIError.StatusCode == 0 {
		cfg.ALIError.StatusCode = ConfigDefault.ALIError.StatusCode
	}
	if cfg.ALIError.Error == "" {
		cfg.ALIError.Error = ConfigDefault.ALIError.Error
	}

	return func(c *fiber.Ctx) error {
		c.Set("X-Authlib-Injector-API-Location", cfg.Location)
		return c.Next()
	}
}
