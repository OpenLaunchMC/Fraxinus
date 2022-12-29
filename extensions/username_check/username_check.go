package usernamecheck

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"

	ferrors "github.com/ProjectOpenLaunch/Fraxinus/errors"
)

type Config struct {
	// Next defines a function to skip this middleware.
	Next func(*fiber.Ctx) bool

	// UsernameCheckRegex is the regex to check the username
	UsernameCheckRegex string

	// UsernameCheckError is the error to return if the username is invalid
	UsernameCheckError ferrors.YggdrasilErrorResponse

	// RegexEngine is the regex engine to use
	RegexEngine func(Regex string , Target string) bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next: func(c *fiber.Ctx) bool {
		json := make(map[string]interface{})
		jsoniter.ConfigFastest.Unmarshal(c.Body(), &json)
		if json["user"] == nil {
			return true
		} else if json["user"].(string) == "" {
			return true
		} else if json["username"].(string) == "" {
			return true
		}
		return false
	},
	UsernameCheckRegex: "[a-zA-Z0-9_]*",
	UsernameCheckError: ferrors.YggdrasilErrorResponse{
		StatusCode: 400,
		Error: "IllegalStateException",
		ErrorMessage: "Invalid characters in username.",
		Cause: "Username contains invalid characters.",
	},
	RegexEngine : func(r string,t string) bool {
		regx , err := regexp.Compile(r)
		if err != nil {
			return false
		}
		result := regx.FindString(t)
		if result == "" {
			return true
		} else {
			return false
		}
	}
}

// New creates a new middleware handler
func New(config ...Config) func(*fiber.Ctx) error {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.UsernameCheckRegex == "" {
		cfg.UsernameCheckRegex = ConfigDefault.UsernameCheckRegex
	}
	if cfg.UsernameCheckError == (ferrors.YggdrasilErrorResponse{}) {
		cfg.UsernameCheckError = ConfigDefault.UsernameCheckError
	}
	if cfg.RegexEngine == nil {
		cfg.RegexEngine = ConfigDefault.RegexEngine
	}
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Skip middleware if Next returns true
		if cfg.Next(c) {
			return c.Next()
		}
		// Check if username is valid
		if cfg.RegexEngine(cfg.UsernameCheckRegex,) {
			return c.Status(cfg.UsernameCheckError.StatusCode).JSON(cfg.UsernameCheckError)
		}
		// Continue stack
		return c.Next()
	}
}