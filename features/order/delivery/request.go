package delivery

import (
	"capstone-alta1/features/order"

	"time"
)

type OrderRequest struct {
	EventName           string    `json:"event_name" form:"event_name"`
	StartDate           time.Time `json:"start_date" form:"start_date"`
	EndDate             time.Time `json:"end_date" form:"end_date"`
	EventLocation       string    `json:"event_location" form:"event_location"`
	EventAddress        string    `json:"event_address" form:"event_address"`
	NotesForPartner     string    `json:"notes_for_partner" form:"notes_for_partner"`
	PaymentMethod       string    `json:"payment_method" form:"payment_method"`
	ServiceID           uint      `json:"service_id" form:"service_id"`
	ServiceAdditionalID uint      `json:"service_additional_id" form:"service_additional_id"`
	Qty                 uint      `json:"qty" form:"qty"`
}

func toCore(input OrderRequest, inputClientID uint) order.Core {
	coreInput := order.Core{
		EventName:       input.EventName,
		StartDate:       input.StartDate,
		EndDate:         input.EndDate,
		EventLocation:   input.EventLocation,
		EventAddress:    input.EventAddress,
		NotesForPartner: input.NotesForPartner,
		PaymentMethod:   input.PaymentMethod,
		ServiceID:       input.ServiceID,
		ClientID:        inputClientID,
	}
	return coreInput
}

func toDetailOrder(input OrderRequest) order.DetailOrder {
	coreInput := order.DetailOrder{
		ServiceAdditionalID: input.ServiceAdditionalID,
		Qty:                 input.Qty,
	}
	return coreInput
}
