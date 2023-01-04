package delivery

import (
	"capstone-alta1/features/service"
	"capstone-alta1/utils/helper"
	"time"
)

type ServiceGetAllResponse struct {
	ID               uint    `json:"id"`
	ServiceName      string  `json:"service_name"`
	ServiceCategory  string  `json:"service_category"`
	ServicePrice     uint    `json:"service_price"`
	AverageRating    float64 `json:"average_rating"`
	ServiceImageFile string  `json:"service_image_file"`
	City             string  `json:"city"`
	PartnerID        uint    `json:"partner_id"`
}

type ServiceGetByIdResponse struct {
	ID                 uint                 `json:"id"`
	ServiceName        string               `json:"service_name"`
	ServiceDescription string               `json:"service_description"`
	ServiceIncluded    string               `json:"service_included"`
	ServiceCategory    string               `json:"service_category"`
	ServicePrice       uint                 `json:"service_price"`
	AverageRating      float64              `json:"average_rating"`
	ServiceImageFile   string               `json:"service_image_file"`
	City               string               `json:"city"`
	Partner            PartnerServiceDetail `json:"partner"`
}

type PartnerServiceDetail struct {
	ID                 uint   `json:"id"`
	CompanyName        string `json:"company_name"`
	CompanyPhone       string `json:"company_phone"`
	CompanyCity        string `json:"company_city"`
	CompanyImageFile   string `json:"company_image_file"`
	CompanyAddress     string `json:"company_address"`
	LinkWebsite        string `json:"link_website"`
	VerificationStatus string `json:"verification_status"`
	UserID             uint   `json:"user_id"`
}

type ServiceAdditionalResponse struct {
	ServiceAdditionalID uint   `json:"service_additional_id"`
	AdditionalName      string `json:"additional_name"`
	AdditionalPrice     uint   `json:"additional_price"`
	ServiceName         string `json:"service_name"`
	ServiceID           uint   `json:"service_id"`
	AdditionalID        uint   `json:"additional_id"`
	PartnerID           uint   `json:"partner_id"`
}

type ServiceReviewResponse struct {
	ID        uint    `json:"id"`
	Review    string  `json:"review"`
	Rating    float64 `json:"rating"`
	OrderID   uint    `json:"order_id"`
	ClientID  uint    `json:"client_id"`
	ServiceID uint    `json:"service_id"`
}

type ServiceDiscussionResponse struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	PartnerID uint      `json:"partner_id"`
	ClientID  uint      `json:"client_id"`
	ServiceID uint      `json:"service_id"`
}

type ServiceAvailabilityResponse struct {
	ServiceName        string `json:"service_name"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	AvailabilityStatus string `json:"availability_status"`
}

func fromCoreGetAll(dataCore service.Core) ServiceGetAllResponse {
	return ServiceGetAllResponse{
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

func fromCoreGetById(dataCore service.ServiceDetailJoinPartner) ServiceGetByIdResponse {
	return ServiceGetByIdResponse{
		ID:                 dataCore.ID,
		ServiceName:        dataCore.ServiceName,
		ServiceDescription: dataCore.ServiceDescription,
		ServiceIncluded:    dataCore.ServiceIncluded,
		ServiceCategory:    dataCore.ServiceCategory,
		ServicePrice:       dataCore.ServicePrice,
		AverageRating:      dataCore.AverageRating,
		ServiceImageFile:   dataCore.ServiceImageFile,
		City:               dataCore.City,
		Partner: PartnerServiceDetail{
			ID:                 dataCore.PartnerID,
			CompanyName:        dataCore.CompanyName,
			CompanyPhone:       dataCore.CompanyPhone,
			CompanyAddress:     dataCore.CompanyAddress,
			CompanyCity:        dataCore.City,
			CompanyImageFile:   dataCore.CompanyImageFile,
			LinkWebsite:        dataCore.LinkWebsite,
			VerificationStatus: dataCore.VerificationStatus,
			UserID:             dataCore.UserID,
		},
	}
}

func fromCoreJoinServiceAdditional(dataCore service.JoinServiceAdditional) ServiceAdditionalResponse {
	return ServiceAdditionalResponse{
		ServiceAdditionalID: dataCore.ServiceAdditionalID,
		AdditionalName:      dataCore.AdditionalName,
		AdditionalPrice:     dataCore.AdditionalPrice,
		ServiceName:         dataCore.ServiceName,
		AdditionalID:        dataCore.AdditionalID,
		ServiceID:           dataCore.ServiceID,
		PartnerID:           dataCore.PartnerID,
	}
}

func fromCoreReview(dataCore service.Review) ServiceReviewResponse {
	return ServiceReviewResponse{
		ID:        dataCore.ID,
		Review:    dataCore.Review,
		Rating:    dataCore.Rating,
		OrderID:   dataCore.OrderID,
		ClientID:  dataCore.ClientID,
		ServiceID: dataCore.ServiceID,
	}
}

func fromCoreDiscussion(dataCore service.Discussion) ServiceDiscussionResponse {
	return ServiceDiscussionResponse{
		ID:        dataCore.ID,
		Comment:   dataCore.Comment,
		CreatedAt: dataCore.CreatedAt,
		ClientID:  dataCore.ClientID,
		PartnerID: dataCore.PartnerID,
		ServiceID: dataCore.ServiceID,
	}
}

func fromCoreAvailability(dataCore service.Order) ServiceAvailabilityResponse {
	return ServiceAvailabilityResponse{
		ServiceName:        dataCore.ServiceName,
		StartDate:          helper.GetDateFormated(dataCore.StartDate),
		EndDate:            helper.GetDateFormated(dataCore.EndDate),
		AvailabilityStatus: dataCore.AvailabilityStatus,
	}
}

func fromCoreList(dataCore []service.Core) []ServiceGetAllResponse {
	var dataResponse []ServiceGetAllResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreGetAll(v))
	}
	return dataResponse
}

func fromCoreListJoinServiceAdditional(dataCore []service.JoinServiceAdditional) []ServiceAdditionalResponse {
	var dataResponse []ServiceAdditionalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreJoinServiceAdditional(v))
	}
	return dataResponse
}

func fromCoreListReview(dataCore []service.Review) []ServiceReviewResponse {
	var dataResponse []ServiceReviewResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreReview(v))
	}
	return dataResponse
}

func fromCoreListDiscussion(dataCore []service.Discussion) []ServiceDiscussionResponse {
	var dataResponse []ServiceDiscussionResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreDiscussion(v))
	}
	return dataResponse
}
