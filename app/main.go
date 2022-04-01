package main

import (
	"order-management/config"
	"order-management/routes"
	"github.com/gofiber/fiber/v2"
)

func main(){
		app := fiber.New()
		config.Connectdb()
		routes.RoutesRegistry(app)
		app.Listen(":3000")
}


