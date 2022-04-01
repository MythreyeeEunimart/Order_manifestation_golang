package models

import (
	"github.com/kamva/mgm/v3"
)

type ErrorMsg struct {
	mgm.DefaultModel `bson:",inline"`
	Status           bool        `json:"status" bson:"status"`
	Message          string      `json:"Message" bson:"Message"`
	ErrorObj         ErrorObject `json:"errorobj" bson:"errorobj"`
}

type ErrorObject struct {
	ErrorCode   string `json:"errorcode" bson:"errorcode"`
	Description string `json:"description" bson:"description"`
}

type ValidationError struct {
	Status   string                `json:"status" bson:"status"`
	Message  string                `json:"Message" bson:"Message"`
	ErrorObj ValidationErrorObject `json:"errorobj" bson:"errorobj"`
}

type ValidationErrorObject struct {
	ErrorCode   string              `json:"errorcode" bson:"errorcode"`
	Description ValidateDescription `json:"description" bson:"description"`
}

type ValidateDescription struct {
	Body []string `json:"body" bson:"body"`
}

type ValidationErrorData struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value interface{} `json:"value"`
}
