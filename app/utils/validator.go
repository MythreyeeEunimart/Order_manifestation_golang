package utils

import (
	"fmt"
	"order-management/models"
	"order-management/services/error_codes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func BodyValidation (c *fiber.Ctx, validationStruct interface{}) error{

    err := c.BodyParser(validationStruct)
    if err != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   err.Error(),
			"errorCode": "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}

    errmsg:= validate.Struct(validationStruct)

    if errmsg != nil{
        var errors []string
        for _,err := range errmsg.(validator.ValidationErrors){
            var response models.ValidationErrorData
            response.Field = err.Field()
            response.Tag = err.Tag()
            response.Value = err.Value()
            fmt.Println("11",response.Value)
            errors = append(errors, response.Tag+" "+response.Field)
        }
        return c.JSON(models.ValidationError{
            Status: "false",
            Message:"Error in input validation",
            ErrorObj : models.ValidationErrorObject{
                ErrorCode : "VDI_ORDERS_VALIDATION",
                Description: models.ValidateDescription{
                    Body: errors,
                },
        },

                })
    }
    return c.Next()
}

func QueryValidation (c *fiber.Ctx, validationStruct interface{}) error{

    err := c.QueryParser(validationStruct)
    if err != nil {
		errResponse := map[string]string{
			"status":    "false",
			"message":   err.Error(),
			"errorCode": "FAILED_TO_BODY-PARSE",
		}
		return error_codes.AppError(c, errResponse)
	}

    errmsg:= validate.Struct(validationStruct)

    if errmsg != nil{
        var errors []string
        for _,err := range errmsg.(validator.ValidationErrors){
            var response models.ValidationErrorData
            response.Field = err.Field()
            response.Tag = err.Tag()
            response.Value = err.Param()
            errors = append(errors, "Query: "+response.Tag+" "+response.Field)
        }
        return c.JSON(models.ValidationError{
            Status: "false",
            Message:"Error in input validation",
            ErrorObj : models.ValidationErrorObject{
                ErrorCode : "VDI_ORDERS_VALIDATION",
                Description: models.ValidateDescription{
                    Body: errors,
                },
        },

                })
    }
    return c.Next()
}