package validation_schemas

type AccountIdSchema struct {
	Account_id string `json:"account_id" bson:"account_id" validate:"required"`
}

type CreateOrderManifestationSchemaData struct {
	Data CreateOrderManifestationSchema `json:"data" bson:"data" validate:"required"`
}

type CreateOrderManifestationSchema struct {
	Account_id           string                `json:"account_id" bson:"account_id" validate:"required"`
	Manifest_date        string                `json:"manifest_date" bson:"manifest_date"`
	Manifestation_number string                `json:"manifestation_number" bson:"manifestation_number" validate:"required"`
	Notes                string                `json:"notes" bson:"notes"`
	Order_count          int                   `json:"order_count" bson:"order_count" validate:"min=1"`
	Value                int                   `json:"value" bson:"value" `
	No_of_boxes          int                   `json:"no_of_boxes" bson:"no_of_boxes"`
	Shipping_partner_id  string                `json:"shipping_partner_id" bson:"shipping_partner_id"`
	Status               string                `json:"status" bson:"status" validate:"required,oneof='new' 'approved'"`
	Attachments          []ValidateAttachments `json:"attachments" bson:"attachments" validate:"min=1,max=5"`
}

type ValidateAttachments struct {
	File_path string `json:"file_path" bson:"file_path" validate:"required"`
	File_name string `json:"file_name" bson:"file_name" validate:"required"`
}

type UpdateOrderManifestationSchemaData struct {
	Data UpdateOrderManifestationSchema `json:"data" bson:"data" validate:"required"`
}

type UpdateOrderManifestationSchema struct {
	Account_id           string                `json:"account_id" bson:"account_id" validate:"required"`
	Manifestation_number string                `json:"manifestation_number" bson:"manifestation_number" validate:"required"`
	Notes                string                `json:"notes" bson:"notes" validate:"required"`
	Value                string                `json:"value" bson:"value" validate:"required"`
	No_of_boxes          string                `json:"no_of_boxes" bson:"no_of_boxes" validate:"required"`
	Shipping_partner_id  string                `json:"shipping_partner_id" bson:"shipping_partner_id" validate:"required"`
	Attachments          []ValidateAttachments `json:"attachments" bson:"attachments" validate:"required,min=1,max=5"`
}

type UpdateStatusOrderManifestationSchemaData struct {
	Data UpdateStatusOrderManifestationSchema `json:"data" bson:"data" validate:"required"`
}

type UpdateStatusOrderManifestationSchema struct {
	Account_id           string `json:"account_id" bson:"account_id" validate:"required"`
	Manifestation_number string `json:"manifestation_number" bson:"manifestation_number" validate:"required"`
	Status               string `json:"status" bson:"status" validate:"required,oneof='new' 'approved'"`
}

type GetOrderManifestationSchema struct {
	Account_id           string `json:"account_id" bson:"account_id" validate:"required"`
	Manifestation_number string `json:"manifestation_number" bson:"manifestation_number" validate:"required"`
}

type GetOrderManifestationListSchema struct {
	Account_id           string   `json:"account_id" bson:"account_id" validate:"required"`
	Manifestation_number string   `json:"manifestation_number" bson:"manifestation_number" validate:"required"`
	Shipping_partner_id  []string `json:"shipping_partner_id" bson:"shipping_partner_id" validate:"required"`
	Status               []string `json:"status" bson:"status" validate:"required"`
}

//how to give array in postman
// how to set oneof value
//confussion in the required field
