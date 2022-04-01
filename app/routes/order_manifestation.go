package routes

import (
	"order-management/utils"
	"order-management/routes/validation_schemas"
	"order-management/services/order_manifestation"
	"github.com/gofiber/fiber/v2"
)


func OrderManifestationRoutes (router fiber.Router){

	router.Post("/create", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.CreateOrderManifestationSchemaData)
        return utils.BodyValidation(c, validation)
    },func(c *fiber.Ctx) error {
        validation := new(validation_schemas.AccountIdSchema)
        return utils.QueryValidation(c, validation)
    },order_manifestation.CreateOrderManifestation )


    router.Post("/update", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.UpdateOrderManifestationSchemaData)
        return utils.BodyValidation(c, validation)
    },func(c *fiber.Ctx) error {
        validation := new(validation_schemas.AccountIdSchema)
        return utils.QueryValidation(c, validation)
    },order_manifestation.UpdateOrderManifestation )


    router.Post("/update_status", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.UpdateStatusOrderManifestationSchemaData)
        return utils.BodyValidation(c, validation)
    },func(c *fiber.Ctx) error {
        validation := new(validation_schemas.AccountIdSchema)
        return utils.QueryValidation(c, validation)
    },order_manifestation.UpdateOrderManifestationStatus )


    router.Post("/get", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.GetOrderManifestationSchema)
        return utils.QueryValidation(c, validation)
    },order_manifestation.GetManifestation )

    router.Post("/list", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.GetOrderManifestationListSchema)
        return utils.QueryValidation(c, validation)
    },order_manifestation.GetManifestation )

}

// app.Post("/create", order_manifestation.CreateOrderManifestation)
// 	app.Post("/update", order_manifestation.UpdateOrderManifestation )
// 	app.Get("/get", order_manifestation.GetManifestation)
// 	app.Get("/list", order_manifestation.ManifestationList)
// 	app.Post("/update_status", order_manifestation.UpdateOrderManifestationStatus)