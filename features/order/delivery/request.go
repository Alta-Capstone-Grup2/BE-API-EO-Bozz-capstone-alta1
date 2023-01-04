package delivery

import (
	"capstone-alta1/features/order"
	"capstone-alta1/utils/helper"
)

type OrderRequest struct {
	EventName       string               `json:"event_name" form:"event_name"`
	StartDate       string               `json:"start_date" form:"start_date"`
	EndDate         string               `json:"end_date" form:"end_date"`
	EventLocation   string               `json:"event_location" form:"event_location"`
	EventAddress    string               `json:"event_address" form:"event_address"`
	NotesForPartner string               `json:"notes_for_partner" form:"notes_for_partner"`
	PaymentMethod   string               `json:"payment_method" form:"payment_method"`
	ServiceID       uint                 `json:"service_id" form:"service_id"`
	OrderDetails    []OrderDetailRequest `json:"order_details" form:"order_details"`
}

// type OrderStatusRequest struct {
// 	OrderStatus string `json:"order_status" form:"order_status"`
// }

type OrderDetailRequest struct {
	ServiceAdditionalID uint `json:"service_additional_id" form:"service_additional_id"`
	Qty                 uint `json:"qty" form:"qty"`
}

func toCore(input OrderRequest, inputClientID uint) order.Core {
	coreInput := order.Core{
		EventName:       input.EventName,
		StartDate:       helper.GetDateTimeFormatedToTime(input.StartDate + " 00:00:00"),
		EndDate:         helper.GetDateTimeFormatedToTime(input.EndDate + " 00:00:00"),
		EventLocation:   input.EventLocation,
		EventAddress:    input.EventAddress,
		NotesForPartner: input.NotesForPartner,
		PaymentMethod:   input.PaymentMethod,
		ServiceID:       input.ServiceID,
		ClientID:        inputClientID,
	}
	return coreInput
}

func toDetailOrder(input OrderDetailRequest) order.DetailOrder {
	coreInput := order.DetailOrder{
		ServiceAdditionalID: input.ServiceAdditionalID,
		Qty:                 input.Qty,
	}
	return coreInput
}

func toDetailOrderList(requestData []OrderDetailRequest) []order.DetailOrder {
	var dataCore []order.DetailOrder
	for _, v := range requestData {
		dataCore = append(dataCore, toDetailOrder(v))
	}
	return dataCore
}

func toCoreStatus(inputCancel string, orderId uint) order.Core {
	coreInput := order.Core{
		ID:          orderId,
		OrderStatus: inputCancel,
	}
	return coreInput
}

type UpdateMidtransRequest struct {
	OrderID     string `json:"order_id" form:"order_id"`
	OrderStatus string `json:"transaction_status" form:"transaction_status"`
}

func toUpdateMidtrans(input UpdateMidtransRequest) order.Core {
	coreInput := order.Core{
		MidtransTransactionID: input.OrderID,
		OrderStatus:           input.OrderStatus,
	}
	return coreInput
}
