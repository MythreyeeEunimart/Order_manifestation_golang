package routes

import (
	// "order-management/services"
	"order-management/services/order_manifestation"
	"github.com/gofiber/fiber/v2"
)


func UserRoutes(app *fiber.App){
	app.Post("/create", order_manifestation.CreateOrderManifestation)
	app.Post("/update", order_manifestation.UpdateOrderManifestation )
	app.Get("/get", order_manifestation.GetManifestation)
}