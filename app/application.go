package app 

import (
	"os"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/postapi/routes"
)


func CreateApp() *fiber.App {
	app := fiber.New()

	os.Setenv("TOEKN_SECRET_KEY", "secret_key_test_me")

	routes.Api(app) 


	return app
}