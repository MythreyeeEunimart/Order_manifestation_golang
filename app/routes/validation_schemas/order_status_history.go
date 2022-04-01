package validation_schemas

import ("time")

type CreateOrderStatusHistorySchemaData struct {
	Data CreateOrderStatusHistorySchema `json:"data" bson:"data" validate:"required"`
}

type CreateOrderStatusHistorySchema struct {
	Account_id          string    `json:"account_id" bson:"account_id" validate:"required"`
	Channel_id          string    `json:"channel_id" bson:"channel_id" `
	Store_id            string    `json:"store_id" bson:"store_id" `
	Order_id            string    `json:"order_id" bson:"order_id" validate:"required"`
	Tracking_number     string    `json:"tracking_number" bson:"tracking_number" `
	Shipping_partner_id string    `json:"shipping_partner_id" bson:"shipping_partner_id" `
	Status              string    `json:"status" bson:"status" validate:"required oneof= 'incomplete' 'new' 'pending' 'processing' 'in_packing' 'packed' 'shipped' 'in_transit' 'delivered' 'delivery_failed' 'short_pick' 'waiting_for_cancellation' 'waiting_for_confirmation' 'cancelled' 'cancellation_failed' 'returned' 'back_order' 'pickup_created' 'pickup_updated' 'address_not_found' 'assignation_failed' 'assignation_pending' 'assignation_failed_volume' 'assignation_failed_weight' 'assignation_failed_shipment' 'assignation_failed_distance' 'assignation_failed_no_quantity' 'assignation_failed_service_type' 'assignation_failed_pickup_point' 'assignation_failed_no_active_warehouse' 'assignation_not_servicable_geo_location' 'assignation_no_active_services_available' 'assignation_no_service_providers_available' 'waiting_for_replenishment' 'waiting_for_return_verification' 'reallocate_service_provider' 'replenishment_created'"`
	Status_timestamp    time.Time `json:"status_timestamp" bson:"status_timestamp" `
	Comments            string    `json:"comments" bson:"comments" `
}

type UpdateOrderStatusHistorySchemaData struct {
	Data UpdateOrderStatusHistorySchema `json:"data" bson:"data" validate:"required"`
}

type UpdateOrderStatusHistorySchema struct{
	Account_id          string    `json:"account_id" bson:"account_id" validate:"required"`
	Channel_id          string    `json:"channel_id" bson:"channel_id" validate:"required"`
	Store_id            string    `json:"store_id" bson:"store_id" validate:"required" `
	Order_id            string    `json:"order_id" bson:"order_id" validate:"required"`
	Tracking_number     string    `json:"tracking_number" bson:"tracking_number" validate:"required"`
	Shipping_partner_id string    `json:"shipping_partner_id" bson:"shipping_partner_id" validate:"required"`
}

type GetOrderStatusHistorySchema struct{
	Account_id          string    `json:"account_id" bson:"account_id" validate:"required"`
	Order_id            string    `json:"order_id" bson:"order_id" validate:"required"`
}