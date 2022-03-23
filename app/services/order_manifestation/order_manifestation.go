package order_manifestation

import (
	"context"
	"fmt"
	// "fmt"
	"order-management/models"
	"order-management/services/error_codes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOrderManifestation(c *fiber.Ctx) error {
    manifestCollection := mgm.CollectionByName("order_manifestation")
	var data models.Data

	err := c.BodyParser(&data)
	if err != nil{

		errResponse := map[string]string{
			"status"    : "false",
			"message"   :  err.Error(),
			"errorCode" :  "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c,errResponse)
    }

	data.Data.Manifest_date = time.Now().Format(time.RFC3339)

	obj := bson.M{
		"manifestation_number" : data.Data.Manifestation_number,
	}
    result := CheckManifestation(c,obj)
	if result == true { 
		_,err:= manifestCollection.InsertOne(context.Background(),data.Data )
		if err != nil{
			errResponse := map[string]string{
				"status"    : "false",
				"message"   :  err.Error(),
				"errorCode" :  "FAILED_TO_CREATE_MANIFESTATION",
			}
			return error_codes.AppError(c,errResponse)

		}else{
			return c.JSON(models.SuccessResponse{
			Status: true,
			Message:"Order manifestation got Created Successfully",
		})
	}
	} else{
		var updateData models.OrderManifestation
		updateData = data.Data
		filter := bson.D{{"manifestation_number" ,data.Data.Manifestation_number }}
		_,errmsg := manifestCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateData})
		if errmsg != nil {
			errResponse := map[string]string{
				"status"    : "false",
				"message"   : errmsg.Error(),
				"errorCode" : "FAILED_TO_UPDATE_MANIFESTATION",
			}
			return error_codes.AppError(c,errResponse)

		}else{
			return c.JSON(models.SuccessResponse{
				Status: true,
				Message:"Manifestation already exist, so updated.",
			})
		}
	}
}

func UpdateOrderManifestation(c *fiber.Ctx) error{
    manifestCollection := mgm.CollectionByName("order_manifestation")
	var data models.Data

	err := c.BodyParser(&data)
	if err != nil{
		errResponse := map[string]string{
			"status"    : "false",
			"message"   :  "",
			"errorCode" :  "FAILED_TO_BODY-PARSE",
			
		}
		return error_codes.AppError(c,errResponse)
    }

	obj := bson.M{
		"manifestation_number" : data.Data.Manifestation_number,
	}
	result := CheckManifestation(c,obj)

	if result == true{
		return c.JSON(models.SuccessResponse{
			Status:   false,
			Message: "Manifestation not found to update.",

		})
	}else{
		var updateData models.OrderManifestation
		updateData = data.Data
		filter := bson.D{{"manifestation_number" ,data.Data.Manifestation_number }}
		_,errmsg := manifestCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateData})
		if errmsg != nil {
			errResponse := map[string]string{
				"status"    : "false",
				"message"   : errmsg.Error(),
				"errorCode" : "FAILED_TO_UPDATE_MANIFESTATION",
			}
			return error_codes.AppError(c,errResponse)

		}else{
			return c.JSON(models.SuccessResponse{
				Status: true,
				Message:"Manifestation already exist, so updated.",
			})
		}
	}
}

func CheckManifestation(c *fiber.Ctx, data map[string]interface{}) bool {
	manifestCollection := mgm.CollectionByName("order_manifestation")
	var details models.Data

	manifestCollection.FindOne(context.Background(),data).Decode(&details.Data)

	if details.Data.Manifestation_number != "" {
		return false
	}else{
		return true
	}
}

func GetManifestation(c *fiber.Ctx) error {

    manifestCollection := mgm.CollectionByName("order_manifestation")
	var details models.OrderManifestation
	 obj := bson.M{
		 "manifestation_number" : c.Query("manifestation_number"),
	 }
	 fmt.Println(obj)

	errmsg := manifestCollection.FindOne(context.Background(),obj).Decode(&details)
	if errmsg != nil {
		errResponse := map[string]string{
			"status"    : "false",
			"message"   :  errmsg.Error(),
			"errorCode" :  "FAILED_TO_FETCH_MANIFESTATION",
		}
		return error_codes.AppError(c,errResponse)
	} 	
	return c.JSON(models.Response{
				Status: true, 
				Data:   details,
	})
}

	


