package models

import (
	"time"
	"github.com/kamva/mgm/v3"
)

type Data struct{
	Data OrderManifestation `json:"data" bson:"data"`
}

type OrderManifestation struct{
	mgm.DefaultModel    `bson:",inline"`
	Manifest_date         string            	`json:"manifest_date" bson:"manifest_date,omitempty"`
	Manifestation_number  string      			`json:"manifestation_number" bson:"manifestation_number,omitempty" validation:"required"`
	Notes                 string                `json:"notes" bson:"notes,omitempty"`
	Order_count           int                   `json:"order_count" bson:"order_count,omitempty"`
	Value                 int                    `json:"value" bson:"value,omitempty"`
	No_of_boxes           int                    `json:"no_of_boxes" bson:"no_of_boxes,omitempty"`
	Shipping_partner_id   string        		`json:"shipping_partner_id" bson:"shipping_partner_id,omitempty"`
	Status                string               `json:"status" bson:"status,omitempty"`
	Status_timestamp      time.Time              `json:"status_timestamp" bson:"status_timestamp"`
	Attachments           []Attachments         `json:"attachments" bson:"attachments,omitempty"`
}


type Attachments struct{
	File_path       string                 `json:"file_path" bson:"file_path,omitempty"`
	File_name      string                  `json:"file_name" bson:"file_name,omitempty"`
}

// type PaginationData struct{
// 	OrderManifestation   []OrderManifestation  `json:"order_manifestation" bson:"order_manifestation"`
//     Pagination           Pagination            `json:"pagination" bson:"pagination"`
// }

type Pagination struct{
	Total_pages    	int      	`json:"total_pages" bson:"total_pages"`
	Per_page       	int			`json:"per_page" bson:"per_page"` 
	Current_page	int			`json:"current_page" bson:"current_page"`
	Next_page		int 		`json:"next_page" bson:"next_page"`
	Pervious_page	int 		`json:"pervious_page" bson:"pervious_page"`
	Total_rows		int 		`json:"total_rows" bson:"Total_rows"`
}

type PaginationResponse struct {
	Status           bool                           `json:"status" bson:"status"`
	Data             []OrderManifestation  			`json:"data" bson:"data"`
    Pagination       Pagination                      `json:"pagination" bson:"pagination"`
	
}

type Response struct{
	
	Status   string        `json:"status" bson:"status"`
    Data     OrderManifestation   `json:"data" bson:"status"`
}

type SuccessResponse struct{
	
	Status   string        `json:"status" bson:"status"`
	Message  string      `json:"Message" bson:"Message"`
}

type GetResponse struct{
	
	Status   string        `json:"status" bson:"status"`
    Data     interface{}  `json:"data" bson:"status"`
}