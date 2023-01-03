package delivery

import (
	"capstone-alta1/features/order"
	"time"
)

type OrderViewResponse struct {
	ID            uint                     `json:"id"`
	EventName     string                   `json:"event_name"`
	StartDate     time.Time                `json:"start_date"`
	EndDate       time.Time                `json:"end_date"`
	EventLocation string                   `json:"event_location"`
	ServiceName   string                   `json:"service_name"`
	GrossAmmount  uint                     `json:"gross_ammount"`
	OrderStatus   string                   `json:"order_status"`
	ServiceID     uint                     `json:"service_id"`
	ClientID      uint                     `json:"client_id"`
	Partner       PartnerOrderViewResponse `json:"partner"`
}

type PartnerOrderViewResponse struct {
	ID                uint   `json:"id"`
	CompanyName       string `json:"company_name"`
	BankName          string `json:"bank_name"`
	BankAccountNumber string `json:"bank_account_number"`
	BankAccountName   string `json:"bank_account_name"`
}

func fromViewCore(dataCore order.OrderJoinPartner) OrderViewResponse {
	return OrderViewResponse{
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
		Partner: PartnerOrderViewResponse{
			ID:                dataCore.PartnerID,
			CompanyName:       dataCore.CompanyName,
			BankName:          dataCore.BankName,
			BankAccountNumber: dataCore.BankAccountNumber,
			BankAccountName:   dataCore.BankAccountName,
		},
	}
}

func fromViewCoreList(dataCore []order.OrderJoinPartner) []OrderViewResponse {
	var dataResponse []OrderViewResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromViewCore(v))
	}
	return dataResponse
}

type OrderResponse struct {
	ID                 uint                  `json:"id"`
	EventName          string                `json:"event_name"`
	StartDate          time.Time             `json:"start_date"`
	EndDate            time.Time             `json:"end_date"`
	StartTime          time.Duration         `json:"start_time"`
	EndTime            time.Duration         `json:"end_time"`
	EventLocation      string                `json:"event_location"`
	EventAddress       string                `json:"event_address"`
	NotesForPartner    string                `json:"notes_for_partner"`
	ServiceName        string                `json:"service_name"`
	ServicePrice       uint                  `json:"service_price"`
	GrossAmmount       uint                  `json:"gross_ammount"`
	PaymentMethod      string                `json:"payment_method"`
	TransactionID      string                `json:"transaction_id"`
	OrderStatus        string                `json:"order_status"`
	VANumber           string                `json:"payment_va_number"`
	PaymentExpiredTime string                `json:"payment_expired_time"`
	PayoutDate         time.Time             `json:"payout_date"`
	PayoutRecieptFile  string                `json:"payout_reciept_file"`
	ServiceID          uint                  `json:"service_id"`
	ClientID           uint                  `json:"client_id"`
	DetailOrders       []DetailOrderResponse `json:"detail_orders"`
}

type DetailOrderResponse struct {
	ID                  uint   `json:"id"`
	AdditionalName      string `json:"additional_name"`
	AdditionalPrice     uint   `json:"additional_price"`
	Qty                 uint   `json:"qty"`
	DetailOrderTotal    uint   `json:"detail_order_total"`
	ServiceAdditionalID uint   `json:"service_additional_id"`
	GrossAmmount        uint   `json:"gross_ammount"`
}

type OrderPaymentResponse struct {
	ID                 uint      `json:"id"`
	EventName          string    `json:"event_name"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	ServiceName        string    `json:"service_name"`
	GrossAmmount       uint      `json:"gross_ammount"`
	PaymentMethod      string    `json:"payment_method"`
	OrderStatus        string    `json:"order_status"`
	TransactionID      string    `json:"transaction_id"`
	VANumber           string    `json:"va_number"`
	PaymentExpiredTime string    `json:"payment_expired_time"`
}

func fromCore(dataCore order.Core, dataCoreDetailOrder []order.DetailOrder) OrderResponse {
	return OrderResponse{
		ID:                 dataCore.ID,
		EventName:          dataCore.EventName,
		StartDate:          dataCore.StartDate,
		EndDate:            dataCore.EndDate,
		StartTime:          dataCore.StartTime,
		EndTime:            dataCore.EndTime,
		EventLocation:      dataCore.EventLocation,
		EventAddress:       dataCore.EventAddress,
		NotesForPartner:    dataCore.NotesForPartner,
		ServiceName:        dataCore.ServiceName,
		ServicePrice:       dataCore.ServicePrice,
		GrossAmmount:       dataCore.GrossAmmount,
		PaymentMethod:      dataCore.PaymentMethod,
		OrderStatus:        dataCore.OrderStatus,
		TransactionID:      dataCore.MidtransTransactionID,
		VANumber:           dataCore.MidtransVaNumber,
		PaymentExpiredTime: dataCore.MidtransExpiredTime,
		PayoutDate:         dataCore.PayoutDate,
		PayoutRecieptFile:  dataCore.PayoutRecieptFile,
		ServiceID:          dataCore.ServiceID,
		ClientID:           dataCore.ClientID,
		DetailOrders:       fromCoreDetailOrderList(dataCoreDetailOrder),
	}
}

func fromCoreDetailOrder(dataCore order.DetailOrder) DetailOrderResponse {
	return DetailOrderResponse{
		ID:                  dataCore.ID,
		AdditionalName:      dataCore.AdditionalName,
		AdditionalPrice:     dataCore.AdditionalPrice,
		Qty:                 dataCore.Qty,
		DetailOrderTotal:    dataCore.DetailOrderTotal,
		ServiceAdditionalID: dataCore.ServiceAdditionalID,
	}
}

func fromCoreDetailOrderList(requestData []order.DetailOrder) []DetailOrderResponse {
	var dataCore []DetailOrderResponse
	for _, v := range requestData {
		dataCore = append(dataCore, fromCoreDetailOrder(v))
	}
	return dataCore
}

func fromCoreToPayment(dataCore order.Core) OrderPaymentResponse {
	return OrderPaymentResponse{
		ID:                 dataCore.ID,
		EventName:          dataCore.EventName,
		StartDate:          dataCore.StartDate,
		EndDate:            dataCore.EndDate,
		ServiceName:        dataCore.ServiceName,
		GrossAmmount:       dataCore.GrossAmmount,
		PaymentMethod:      dataCore.PaymentMethod,
		OrderStatus:        dataCore.OrderStatus,
		TransactionID:      dataCore.MidtransTransactionID,
		VANumber:           dataCore.MidtransVaNumber,
		PaymentExpiredTime: dataCore.MidtransExpiredTime,
	}
}
