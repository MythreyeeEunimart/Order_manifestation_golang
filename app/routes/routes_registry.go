package routes

import "github.com/gofiber/fiber/v2"

const BASE_PATH = "/api/v2/order_management"

func RoutesRegistry(app *fiber.App) {

    base_path := app.Group(BASE_PATH)

    order_manifestation_route := base_path.Group("/order_manifestation")
    OrderManifestationRoutes(order_manifestation_route)

    order_status_history_route := base_path.Group("/order_status_history")
    OrderStatusHisoryRoutes(order_status_history_route)
}

