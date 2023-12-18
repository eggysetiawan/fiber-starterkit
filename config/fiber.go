package config

import (
	"github.com/gofiber/fiber/v2"
)

func RunServe() {
	// Custom config
	_ = fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Api-Ejol",
		AppName:       "Api Ejol v1.0.0",
	})
}
