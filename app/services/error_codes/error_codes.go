package error_codes

import (
	"order-management/models"

	"github.com/gofiber/fiber/v2"
)

func AppError(c *fiber.Ctx, data map[string]string) (error){

	var errMsg = map[string]map[string]string{}

	errMsg["FAILED_TO_CREATE"] = map[string]string{}
    errMsg["FAILED_TO_CREATE"]["errorCode"] = "VDI_SLSA_PRODUCTS_0001"
    errMsg["FAILED_TO_CREATE"]["description"] = "DB error or order_id might be duplicate"

	errMsg["FAILED_TO_FETCH_MANIFESTATION"] = map[string]string{}
    errMsg["FAILED_TO_FETCH_MANIFESTATION"]["errorCode"] = "VDI_SLSA_ORDERS_0017"
    errMsg["FAILED_TO_FETCH_MANIFESTATION"]["description"] = "Failed to get the manifestation data."

    errMsg["FAILED_TO_BODY-PARSE"] = map[string]string{}
    errMsg["FAILED_TO_BODY-PARSE"]["errorCode"] = "NODE_INT_000"
    errMsg["FAILED_TO_BODY-PARSE"]["description"] = "nternal Error, Unable to process the request"

	errMsg["FAILED_TO_UPDATE_MANIFESTATION"] = map[string]string{}
    errMsg["FAILED_TO_UPDATE_MANIFESTATION"]["errorCode"] = "VDI_SLSA_ORDERS_0016"
    errMsg["FAILED_TO_UPDATE_MANIFESTATION"]["description"] = "Unable to update this manifestation."

    errMsg["FAILED_TO_CREATE_MANIFESTATION"] = map[string]string{}
    errMsg["FAILED_TO_CREATE_MANIFESTATION"]["errorCode"] = "VDI_SLSA_ORDERS_0015"
    errMsg["FAILED_TO_CREATE_MANIFESTATION"]["description"] = "Can't create manifestation for this order."



	var errorSchema models.ErrorMsg

	if data["status"] == "false" {
		errorSchema.Status = false
	}else{
		errorSchema.Status =true
	}
	if data["message"] == "" {
		errorSchema.Message = "Something went wrong in Server"
	}else{
		errorSchema.Message = data["message"]
	}
	if data["errorCode"] != "" {
		errorSchema.ErrorObj.ErrorCode = errMsg[data["errorCode"]]["errorCode"]
		errorSchema.ErrorObj.Description = errMsg[data["errorCode"]]["description"]
	}

	return c.JSON(errorSchema)
}

