package delivery

import (
	"capstone-alta1/features/partner"
	"capstone-alta1/utils/helper"
)

type PartnerResponse struct {
	ID                 uint   `json:"id"`
	PICName            string `json:"pic_name"`
	Email              string `json:"email"`
	Role               string `json:"role"`
	PICPosition        string `json:"pic_position"`
	PICPhone           string `json:"pic_phone"`
	PICAddress         string `json:"pic_address"`
	CompanyName        string `json:"company_name"`
	CompanyPhone       string `json:"company_phone"`
	CompanyCity        string `json:"company_city"`
	CompanyImageFile   string `json:"company_image_file"`
	CompanyAddress     string `json:"company_address"`
	LinkWebsite        string `json:"link_website"`
	NIBNumber          string `json:"nib_number"`
	NIBImageFile       string `json:"nib_image_file"`
	SIUPNumber         string `json:"siup_number"`
	SIUPImageFile      string `json:"siup_image_file"`
	Event1Name         string `json:"event1_name"`
	Event1ImageFile    string `json:"event1_image_file"`
	Event2Name         string `json:"event2_name"`
	Event2ImageFile    string `json:"event2_image_file"`
	Event3Name         string `json:"event3_name"`
	Event3ImageFile    string `json:"event4_image_file"`
	BankName           string `json:"bank_name"`
	BankAccountNumber  string `json:"bank_account_number"`
	BankAccountName    string `json:"bank_account_name"`
	VerificationStatus string `json:"verification_status"`
	VerificationLog    string `json:"verification_log"`
	UserID             uint   `json:"user_id"`
}

type PartnerListResponse struct {
	ID                 uint   `json:"id"`
	CompanyName        string `json:"company_name"`
	PICName            string `json:"pic_name"`
	VerificationStatus string `json:"verification_status"`
	RegisterDate       string `json:"register_date"`
	UserID             uint   `json:"user_id"`
}

type ServiceResponse struct {
	ID               uint    `json:"id"`
	ServiceName      string  `json:"service_name"`
	ServiceCategory  string  `json:"service_category"`
	ServicePrice     uint    `json:"service_price"`
	AverageRating    float64 `json:"average_rating"`
	ServiceImageFile string  `json:"service_image_file"`
	City             string  `json:"city"`
	PartnerID        uint    `json:"partner_id"`
}

type OrderResponse struct {
	ID               uint   `json:"id"`
	EventName        string `json:"event_name"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	EventLocation    string `json:"event_location"`
	ServiceName      string `json:"service_name"`
	GrossAmmount     uint   `json:"gross_ammount"`
	OrderStatus      string `json:"order_status"`
	PayoutRecieptUrl string `json:"payout_receipt_url"`
	ServiceID        uint   `json:"service_id"`
	ClientID         uint   `json:"client_id"`
}

type AdditionalResponse struct {
	ID              uint   `json:"id"`
	AdditionalName  string `json:"additional_name"`
	AdditionalPrice uint   `json:"additional_price"`
	PartnerID       uint   `json:"partner_id"`
}

func fromCore(dataCore partner.Core) PartnerResponse {
	return PartnerResponse{
		ID:                 dataCore.ID,
		PICName:            dataCore.User.Name,
		Email:              dataCore.User.Email,
		Role:               dataCore.User.Role,
		PICPosition:        dataCore.PICPosition,
		PICPhone:           dataCore.PICPhone,
		PICAddress:         dataCore.PICAddress,
		CompanyName:        dataCore.CompanyName,
		CompanyPhone:       dataCore.CompanyPhone,
		CompanyCity:        dataCore.CompanyCity,
		CompanyImageFile:   dataCore.CompanyImageFile,
		CompanyAddress:     dataCore.CompanyAddress,
		LinkWebsite:        dataCore.LinkWebsite,
		NIBNumber:          dataCore.NIBNumber,
		NIBImageFile:       dataCore.NIBImageFile,
		SIUPNumber:         dataCore.SIUPNumber,
		SIUPImageFile:      dataCore.SIUPImageFile,
		Event1Name:         dataCore.Event1Name,
		Event1ImageFile:    dataCore.Event1ImageFile,
		Event2Name:         dataCore.Event2Name,
		Event2ImageFile:    dataCore.Event2ImageFile,
		Event3Name:         dataCore.Event3Name,
		Event3ImageFile:    dataCore.Event3ImageFile,
		BankName:           dataCore.BankName,
		BankAccountNumber:  dataCore.BankAccountNumber,
		BankAccountName:    dataCore.BankName,
		VerificationStatus: dataCore.VerificationStatus,
		VerificationLog:    dataCore.VerificationLog,
		UserID:             dataCore.UserID,
	}
}

func fromListCore(dataCore partner.Core) PartnerListResponse {
	return PartnerListResponse{
		ID:                 dataCore.ID,
		CompanyName:        dataCore.CompanyName,
		PICName:            dataCore.User.Name,
		RegisterDate:       helper.GetDateTimeFormated(dataCore.CreatedAt),
		VerificationStatus: dataCore.VerificationStatus,
		UserID:             dataCore.UserID,
	}
}

// func fromCoreList(dataCore []partner.Core) []PartnerResponse {
// 	var dataResponse []PartnerResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCore(v))
// 	}
// 	return dataResponse
// }

func fromListCoreList(dataCore []partner.Core) []PartnerListResponse {
	var dataResponse []PartnerListResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromListCore(v))
	}
	return dataResponse
}

func fromOrderCore(dataCore partner.OrderCore) OrderResponse {
	return OrderResponse{
		ID:               dataCore.ID,
		EventName:        dataCore.EventName,
		StartDate:        helper.GetDateFormated(dataCore.StartDate),
		EndDate:          helper.GetDateFormated(dataCore.EndDate),
		EventLocation:    dataCore.EventLocation,
		ServiceName:      dataCore.ServiceName,
		GrossAmmount:     dataCore.GrossAmmount,
		OrderStatus:      dataCore.OrderStatus,
		PayoutRecieptUrl: dataCore.PayoutRecieptFile,
		ServiceID:        dataCore.ServiceID,
		ClientID:         dataCore.ClientID,
	}
}

func fromOrderCoreList(dataCore []partner.OrderCore) []OrderResponse {
	var dataResponse []OrderResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromOrderCore(v))
	}
	return dataResponse
}

func fromCoreService(dataCore partner.ServiceCore) ServiceResponse {
	return ServiceResponse{
		ID:               dataCore.ID,
		ServiceName:      dataCore.ServiceName,
		ServiceCategory:  dataCore.ServiceCategory,
		ServicePrice:     dataCore.ServicePrice,
		AverageRating:    dataCore.AverageRating,
		ServiceImageFile: dataCore.ServiceImageFile,
		City:             dataCore.City,
		PartnerID:        dataCore.PartnerID,
	}
}

func fromCoreServiceList(dataCore []partner.ServiceCore) []ServiceResponse {
	var dataResponse []ServiceResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreService(v))
	}
	return dataResponse
}

func fromAdditonalCore(dataCore partner.AdditionalCore) AdditionalResponse {
	return AdditionalResponse{
		ID:              dataCore.ID,
		AdditionalName:  dataCore.AdditionalName,
		AdditionalPrice: dataCore.AdditionalPrice,
		PartnerID:       dataCore.PartnerID,
	}
}

func fromAdditionalCoreList(dataCore []partner.AdditionalCore) []AdditionalResponse {
	var dataResponse []AdditionalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromAdditonalCore(v))
	}
	return dataResponse
}
