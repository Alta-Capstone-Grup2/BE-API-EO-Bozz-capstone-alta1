package delivery

import (
	"capstone-alta1/features/order"
	"time"
)

type OrderResponse struct {
	ID            uint      `json:"id"`
	EventName     string    `json:"event_name"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	EventLocation string    `json:"event_location"`
	ServiceName   string    `json:"service_name"`
	GrossAmmount  uint      `json:"gross_ammount"`
	OrderStatus   string    `json:"order_status"`
	ServiceID     uint      `json:"service_id"`
	ClientID      uint      `json:"client_id"`
}

func fromCore(dataCore order.Core) OrderResponse {
	return OrderResponse{
		ID:            dataCore.ID,
		EventName:     dataCore.EventName,
		StartDate:     dataCore.StartDate,
		EndDate:       dataCore.EndDate,
		EventLocation: dataCore.EventLocation,
		ServiceName:   dataCore.ServiceName,
		GrossAmmount:  dataCore.GrossAmmount,
		OrderStatus:   dataCore.OrderStatus,
		ServiceID:     dataCore.ServiceID,
		ClientID:      dataCore.ClientID,
	}
}

func fromCoreList(dataCore []order.Core) []OrderResponse {
	var dataResponse []OrderResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

type OrderDetailResponse struct {
	ID                  uint          `json:"id"`
	EventName           string        `json:"event_name"`
	StartDate           time.Time     `json:"start_date"`
	EndDate             time.Time     `json:"end_date"`
	StartTime           time.Duration `json:"start_time"`
	EndTime             time.Duration `json:"end_time"`
	EventLocation       string        `json:"event_location"`
	EventAddress        string        `json:"event_address"`
	NotesForPartner     string        `json:"notes_for_partner"`
	ServiceName         string        `json:"service_name"`
	ServicePrice        uint          `json:"service_price"`
	GrossAmmount        uint          `json:"gross_ammount"`
	PaymentMethod       string        `json:"payment_method"`
	OrderStatus         string        `json:"order_status"`
	PayoutDate          time.Time     `json:"payout_date"`
	PayoutRecieptFile   string        `json:"payout_reciept_file"`
	ServiceID           uint          `json:"service_id"`
	ClientID            uint          `json:"client_id"`
	ServiceAdditionalID uint          `json:"service_additional_id"`
	AdditionalName      string        `json:"additional_name"`
	AdditionalPrice     uint          `json:"additional_price"`
	Qty                 uint          `json:"qty"`
	DetailOrderTotal    uint          `json:"detail_order_total"`
	OrderID             uint          `json:"order_id"`
}

func fromCoreDetail(dataCore order.Core, dataDetail order.DetailOrder) OrderDetailResponse {
	return OrderDetailResponse{
		ID:                  dataCore.ID,
		EventName:           dataCore.EventName,
		StartDate:           dataCore.StartDate,
		EndDate:             dataCore.EndDate,
		StartTime:           dataCore.StartTime,
		EndTime:             dataCore.EndTime,
		EventLocation:       dataCore.EventLocation,
		EventAddress:        dataCore.EventAddress,
		NotesForPartner:     dataCore.NotesForPartner,
		ServiceName:         dataCore.ServiceName,
		ServicePrice:        dataCore.ServicePrice,
		GrossAmmount:        dataCore.GrossAmmount,
		PaymentMethod:       dataCore.PaymentMethod,
		OrderStatus:         dataCore.OrderStatus,
		PayoutDate:          dataCore.PayoutDate,
		PayoutRecieptFile:   dataCore.PayoutRecieptFile,
		ServiceID:           dataCore.ServiceID,
		ClientID:            dataCore.ClientID,
		ServiceAdditionalID: dataDetail.ServiceAdditionalID,
		AdditionalName:      dataDetail.AdditionalName,
		AdditionalPrice:     dataDetail.AdditionalPrice,
		Qty:                 dataDetail.Qty,
		DetailOrderTotal:    dataDetail.DetailOrderTotal,
		OrderID:             dataDetail.OrderID,
	}
}
