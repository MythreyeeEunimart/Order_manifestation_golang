package order_manifestation

import (
	"context"
	"fmt"
	"order-management/models"
	"order-management/services/error_codes"
	"time"
    "strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"github.com/iancoleman/strcase"
)

func CreateOrderManifestation(c *fiber.Ctx) error {
	manifestCollection := mgm.CollectionByName("order_manifestation")
	var data models.Data

	err := c.BodyParser(&data)
	if err != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   err.Error(),
			"errorCode": "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}

	data.Data.Manifest_date = time.Now().Format(time.RFC3339)

	obj := bson.M{
		"manifestation_number": data.Data.Manifestation_number,
	}
	result := CheckManifestation(c, obj)
	if result.Status == "true" {
		data.Data.Order_count = 1

		err := manifestCollection.Create(&data.Data)
		if err != nil {
			errResponse := map[string]string{
				"status":    "false",
				"message":   err.Error(),
				"errorCode": "FAILED_TO_CREATE_MANIFESTATION",
			}
			return error_codes.AppError(c, errResponse)

		} else {
			return c.JSON(models.SuccessResponse{
				Status:  "true",
				Message: "Order manifestation got Created Successfully",
			})
		}
	} else {
		data.Data.Order_count = result.Data.Order_count+1

		result1 := UpdateOrderManifestationData(c, data, result.Data)

		if result1.Status == "true" {
			return c.JSON(models.SuccessResponse{
				Status:  "true",
				Message: "Manifestation already exist, so updated.",
			})
		} else {
			errResponse := map[string]string{
				"status":    "false",
				"message":   "",
				"errorCode": "FAILED_TO_UPDATE_MANIFESTATION",
			}
			return error_codes.AppError(c, errResponse)
		}
	}
}

func UpdateOrderManifestation(c *fiber.Ctx) error{
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

	if result.Status == "true"{
		return c.JSON(models.SuccessResponse{
			Status:   "false",
			Message: "Manifestation not found to update.",

		})
	}else {

		result1 := UpdateOrderManifestationData(c, data, result.Data)

		if result1.Status == "true" {
			return c.JSON(models.SuccessResponse{
				Status:  "true",
				Message: "Manifestation updated successfully.",
			})
		} else {
			errResponse := map[string]string{
				"status":    "false",
				"message":   "",
				"errorCode": "FAILED_TO_UPDATE_MANIFESTATION",
			}
			return error_codes.AppError(c, errResponse)
		}
	}
}

func CheckManifestation(c *fiber.Ctx, data map[string]interface{}) models.Response {
	manifestCollection := mgm.CollectionByName("order_manifestation")
	var details models.OrderManifestation
	fmt.Println("...",data)
	
	errmsg := manifestCollection.FindOne(context.Background(), data).Decode(&details)
	fmt.Println(details)
	if errmsg != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   errmsg.Error(),
			"errorCode": "FAILED_TO_FETCH_MANIFESTATION",
		}
		error_codes.AppError(c, errResponse)
	}
	if details.Manifestation_number == "" {
		return (models.Response{
			Status: "true",
			Data:   details,
		})
	} else {
		return (models.Response{
			Status: "false",
			Data:   details,
		})

	}
}

func UpdateOrderManifestationData(c *fiber.Ctx, data models.Data, details models.OrderManifestation) models.SuccessResponse {

	manifestCollection := mgm.CollectionByName("order_manifestation")
	data.Data.Value = data.Data.Value + details.Value
	data.Data.No_of_boxes = data.Data.No_of_boxes + details.No_of_boxes
	data.Data.Status_timestamp = time.Now()
	fmt.Println(details.CreatedAt)
	data.Data.CreatedAt = details.CreatedAt

	if details.Status == "new" {
		data.Data.Status = "approved"
	}
	var updateData models.OrderManifestation
	var result models.OrderManifestation
	updateData = data.Data
	fmt.Println("hii",updateData.Order_count)
	details.Attachments = data.Data.Attachments
	filter := bson.M{"manifestation_number": data.Data.Manifestation_number}
	updateData.UpdatedAt = time.Now()

	err := manifestCollection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set":updateData}).Decode(&result)
	if err != nil { 
		errResponse := map[string]string{
			"status":    "false",
			"message":   "22",
			"errorCode": "FAILED_TO_UPDATE_MANIFESTATION",
		}
		error_codes.AppError(c, errResponse)
		return (models.SuccessResponse{
			Status:  "false",
			Message: "Order Manifestation cannot update",
		})

	} else {
		return (models.SuccessResponse{
			Status:  "true",
			Message: "Order Manifestation Updated Successfully",
		})
	}
}

func GetManifestation(c *fiber.Ctx) error {

    manifestCollection := mgm.CollectionByName("order_manifestation")
	var details models.OrderManifestation

	 obj := bson.M{
		 "manifestation_number" : c.Query("manifestation_number"),
	 }

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
				Status: "true",
				Data:   details,
	})
}

func  ManifestationList(c *fiber.Ctx) error{

	var data []models.OrderManifestation
	var pagination models.Pagination
	manifestation := mgm.CollectionByName("order_manifestation")

	find_option := options.Find()
	sort_col := c.Query("sortCol")
    sort_order,_ := strconv.Atoi(c.Query("sortOrder"))

	page_no, _ := strconv.Atoi(c.Query("pageNo","1"))
    per_page, _ := strconv.Atoi(c.Query("perPage", "10"))

	find_obj := bson.M{}

	if sort_order != 0 && sort_col != "" {
        find_option.SetSort(bson.D{{Key:sort_col,Value:sort_order}})
    }

	if c.Query("account_id") != "" {
        find_obj["account_id"] = c.Query("account_id")
    }

	if c.Query("manifestation_number") != "" {
        find_obj["manifestation_number"] = c.Query("manifestation_number")
    }

	if c.Query("status") != "" {
        find_obj["status"] = c.Query("status")
    }

	if c.Query("shipping_partner_id") != "" {
        find_obj["shipping_partner_id"] = c.Query("shipping_partner_id")
    }

	pagination.Current_page = page_no
	pagination.Per_page = per_page

	find_option.SetSkip((int64(page_no-1))*int64(per_page))
    find_option.SetLimit(int64(per_page))

	total_data, _ := manifestation.EstimatedDocumentCount(context.Background())
    pagination.Total_rows = int(total_data)


	total_pages := int(math.Ceil(float64(pagination.Total_rows)/float64(pagination.Per_page)))
    pagination.Total_pages = total_pages

	pagination.Pervious_page = pagination.Current_page -1
	pagination.Next_page = pagination.Current_page +1

	err := manifestation.SimpleFind(&data,find_obj,find_option)

    if err!= nil {
        errResponse := map[string]string{
            "status":"false",
            "message": err.Error(),
            "error_code": "FAILED_TO_FETCH_DATA",
        }
		return error_codes.AppError(c,errResponse)
    }else{
		return c.JSON(models.PaginationResponse{
			Status : true,
			Data: data,
			Pagination: pagination,
		})

	}
}

func UpdateOrderManifestationStatus(c *fiber.Ctx) error {
	manifestCollection := mgm.CollectionByName("order_manifestation")
	var data models.Data
	var result models.OrderManifestation
	
	err := c.BodyParser(&data)
	if err != nil{
		errResponse:= map[string]string{
			"status" :    "false",
			"message" :   "",
			"errorCode" : "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}

	status := data.Data.Status
	camalizeStatus := strcase.ToLowerCamel(status)

	filter:= bson.M{
		"manifestation_number" : data.Data.Manifestation_number,
	}

	update := bson.M{
		"status" : camalizeStatus,
	}

	errmsg := manifestCollection.FindOneAndUpdate(context.Background(),filter, bson.M{"$set" : update}).Decode(&result)
	if errmsg != nil{
		errResponse := map[string]string{
			"status" :"false",
			"message" : "",
			"errorCode" : "FAILED_TO_UPDATE_ORDER",

		}
		return error_codes.AppError(c,errResponse)
	}else{
		return c.JSON(models.SuccessResponse{
			Status:  "true",
			Message: "Order Manifestation Status Updated Successfully.",
		})
	}
}