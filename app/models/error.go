package models

import (
	"github.com/kamva/mgm/v3"
)

type ErrorMsg struct{

    mgm.DefaultModel    `bson:",inline"`
	Status    bool          `json:"status" bson:"status" validate:"required"`
	Message   string        `json:"Message" bson:"Message" validate:"required"`
	ErrorObj  ErrorObject   `json:"errorobj" bson:"errorobj" validate:"required"`
}


type ErrorObject struct{

	ErrorCode    string       `json:"errorcode" bson:"errorcode" validate:"required"`
	Description  string       `json:"description" bson:"description" validate:"required"`
}

