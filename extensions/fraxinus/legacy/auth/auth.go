package legacyauth

import (
	"github.com/gofiber/fiber/v2"

	ferrors "github.com/ProjectOpenLaunch/Fraxinus/errors"
)

type Config struct {
	CheckHeader bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{

}


// New creates a new middleware handler
func New(config ...Config) fiber.Handler {

}
