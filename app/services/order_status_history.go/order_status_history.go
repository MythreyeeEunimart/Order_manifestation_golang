package order_status_history

import (
	"context"
	"order-management/models"
	"order-management/services/error_codes"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOrderStatusHistory(c *fiber.Ctx) error {

	var data models.OrderStatusHistoryData
	statusHistoryCollection := mgm.CollectionByName("order_status_history")

	err := c.BodyParser(&data)
	if err != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   err.Error(),
			"errorcode": "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}

	errmsg := statusHistoryCollection.Create(&data.Data)
	if errmsg != nil {
		errResponse := map[string]string{
			"status":    "false",
			"messgae":   errmsg.Error(),
			"errorCode": "FAILED_TO_CREATE_STATUS_HISTORY",
		}
		return error_codes.AppError(c, errResponse)
	} else {
		return c.JSON(models.SuccessResponse{
			Status:  "true",
			Message: "Order Status History got Created Successfully",
		})
	}
}

func UpdateOrderStatusHistory(c *fiber.Ctx) error {

	var data models.OrderStatusHistoryData
	var result models.OrderStatusHistory
	statusHistoryCollection := mgm.CollectionByName("order_status_history")

	err := c.BodyParser(&data)
	if err != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   err.Error(),
			"errorcode": "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}
	updateData := bson.M{
		"comments": data.Data.Comments,
	}
	findObj := bson.M{
		"order_id":            data.Data.Order_id,
		"channel_id":          data.Data.Channel_id,
		"account_id":          data.Data.Account_id,
		"store_id":            data.Data.Store_id,
		"tracking_number":     data.Data.Tracking_number,
		"shipping_partner_id": data.Data.Shipping_patner_id,
	}

	errmsg := statusHistoryCollection.FindOneAndUpdate(context.Background(), findObj, bson.M{"$set": updateData}).Decode(&result)
	if errmsg != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   errmsg.Error(),
			"errorCode": "FAILED_TO_UPDATE_ORDER",
		}
		return error_codes.AppError(c, errResponse)

	} else {
		return c.JSON(models.SuccessResponse{
			Status:  "true",
			Message: "Order Status History Updated Successfully",
		})
	}
}

func GetOrderStatusHistory(c *fiber.Ctx) error {

	var orderStatus []models.OrderStatusHistory
	statusHistoryCollection := mgm.CollectionByName("order_status_history")

	object := bson.M{
		"account_id": c.Query("account_id"),
		"order_id":   c.Query("order_id"),
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", 1}})

	err := statusHistoryCollection.SimpleFind(&orderStatus, object, findOptions)
	if err != nil {
		errResponse := map[string]string{
			"status":     "false",
			"message":    err.Error(),
			"error_code": "FAILED_TO_FETCH_DATA",
		}
		return error_codes.AppError(c,errResponse)
	}else{
		return c.JSON(models.GetResponse{
			Status: "true" ,
			Data : orderStatus, 
		})
	}

}