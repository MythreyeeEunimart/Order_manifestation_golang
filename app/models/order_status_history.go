package models

import(
	"github.com/kamva/mgm/v3"
	"time"
)

type OrderStatusHistoryData struct{

	Data OrderStatusHistory  `json:"data" bson:"data"`
}
 

type OrderStatusHistory struct{
	mgm.DefaultModel `  bson:",inline"`
	Account_id  		string            `json:"account_id" bson:"account_id" validate:"required"`
	Channel_id  		string            `json:"channel_id" bson:"channel_id" validate:"required"`
	Store_id    		string            `json:"store_id" bson:"store_id" validate:"required"`  
	Order_id    		string            `json:"order_id" bson:"order_id" validate:"required"`
	Tracking_number 	string        	  `json:"tracking_number" bson:"tracking_number" validate:"required"`
	Shipping_patner_id 	string     		  `json:"shipping_patner_id" bson:"shipping_patner_id" validate:"required"`
	Status 				string            `json:"status" bson:"status" validate:"required"`
	Status_timestamp 	time.Time      	  `json:"status_timestamp" bson:"status_timestamp" validate:"required"`
	Comments 			string            `json:"comments" bson:"comments" validate:"required"`
}