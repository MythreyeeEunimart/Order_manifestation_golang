package models

import (
	"github.com/kamva/mgm/v3"
)

type Data struct{
	Data OrderManifestation `json:"data" bson:"data"`
}

type OrderManifestation struct{
	mgm.DefaultModel    `bson:",inline"`
	Manifest_date         string            	`json:"manifest_date" bson:"manifest_date"`
	Manifestation_number  string      			`json:"manifestation_number" bson:"manifestation_number" validation:"required"`
	Notes                 string                `json:"notes" bson:"notes"`
	Order_count           string                `json:"order_count" bson:"order_count"`
	Value                 string                `json:"value" bson:"value"`
	No_of_boxes           string                `json:"no_of_boxes" bson:"no_of_boxes"`
	Shipping_partner_id   string        		`json:"shipping_partner_id" bson:"shipping_partner_id"`
	Status                string               `json:"status" bson:"status"`
	Status_timestamp      string               `json:"status_timestamp" bson:"status_timestamp"`
	Attachments           []Attachments         `json:"attachments" bson:"attachments"`
}


type Attachments struct{
	 
	File_path       string                 `json:"file_path" bson:"file_path"`
	File_name      string                  `json:"file_name" bson:"file_name"`

}

type Response struct{
	
	Status   bool        `json:"status" bson:"status"`
    Data     OrderManifestation   `json:"data" bson:"status"`
}

type SuccessResponse struct{
	
	Status   bool        `json:"status" bson:"status"`
	Message  string      `json:"Message" bson:"Message"`
}