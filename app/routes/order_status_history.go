package routes

import(
	"github.com/gofiber/fiber/v2"
	"order-management/services/order_status_history.go"
	"order-management/routes/validation_schemas"
	"order-management/utils"
)

func OrderStatusHisoryRoutes (router fiber.Router){

	router.Post("/create", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.CreateOrderStatusHistorySchemaData)
        return utils.BodyValidation(c, validation)
    },func(c *fiber.Ctx) error {
        validation := new(validation_schemas.AccountIdSchema)
        return utils.QueryValidation(c, validation)
    },order_status_history.CreateOrderStatusHistory )
	

	router.Post("/update", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.UpdateOrderStatusHistorySchemaData)
        return utils.BodyValidation(c, validation)
    },func(c *fiber.Ctx) error {
        validation := new(validation_schemas.AccountIdSchema)
        return utils.QueryValidation(c, validation)
    },order_status_history.UpdateOrderStatusHistory )


	router.Post("/get", func(c *fiber.Ctx) error {
        validation := new(validation_schemas.GetOrderStatusHistorySchema)
        return utils.QueryValidation(c, validation)
    },order_status_history.GetOrderStatusHistory )

}
