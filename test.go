package services

import (
	// "fmt"
	"context"
	// "fmt"
	"order-management/models"
	"order-management/services/error_codes"

	// "github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"

	// "github.com/rbrahul/gofp"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOrderStatusHistory(c *fiber.Ctx) error{

	var data models.OrderStatusHistoryData
	statusHistoryCollection := mgm.CollectionByName("order_status_history")
	

	err:= c.BodyParser(&data)
	if err != nil{
		errResponse := map[string]string{
			"status" :"false",
			"message" : err.Error(),
			"errorcode" : "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c,errResponse)
	}

	errmsg := statusHistoryCollection.Create(&data.Data)
	if errmsg != nil {
		errResponse := map[string]string{
			"status" :"false",
			"messgae" : errmsg.Error(),
			"errorCode" :"FAILED_TO_CREATE_STATUS_HISTORY",
		}
		return error_codes.AppError(c,errResponse)
	}else{
		return c.JSON(models.SuccessResponse{
			Status: "true",
			Message: "Order Status History got Created Successfully",
		})
	}
}

func UpdateOrderStatusHistory(c *fiber.Ctx) error{

	var data models.OrderStatusHistoryData
	var result models.OrderStatusHistory
	statusHistoryCollection := mgm.CollectionByName("order_status_history")
	
	err:= c.BodyParser(&data)
	if err != nil{
		errResponse := map[string]string{
			"status" :"false",
			"message" : err.Error(),
			"errorcode" : "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c,errResponse)
	}
	updateData := bson.M{
		"comments" : data.Data.Comments,
	}
	findObj := bson.M{
		"order_id" :data.Data.Order_id,
		"channel_id" : data.Data.Channel_id,
		"account_id" : data.Data.Account_id,
		"store_id" : data.Data.Store_id,
		"tracking_number" : data.Data.Tracking_number,
		"shipping_partner_id" : data.Data.Shipping_patner_id,

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
			Status: "true",
			Message: "Order Status History Updated Successfully",
		})
	}
}
