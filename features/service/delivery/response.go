package delivery

import "capstone-alta1/features/service"

type ServiceResponse struct {
	ID              uint    `json:"id"`
	ServiceName     string  `json:"service_name"`
	ServiceCategory string  `json:"service_category"`
	ServicePrice    uint    `json:"service_price"`
	AverageRating   float64 `json:"average_rating"`
	ServiceImageUrl string  `json:"service_image_file"`
	City            string  `json:"city"`
	PartnerID       uint    `json:"partner_id"`
}

type ServiceAdditionalResponse struct {
	ID              uint   `json:"id"`
	AdditionalName  string `json:"additional_name"`
	AdditionalPrice uint   `json:"additional_price"`
	PartnerID       uint   `json:"partner_id"`
	ServiceID       uint   `json:"service_id"`
}

func fromCore(dataCore service.Core) ServiceResponse {
	return ServiceResponse{
		ID:              dataCore.ID,
		ServiceName:     dataCore.ServiceName,
		ServiceCategory: dataCore.ServiceCategory,
		ServicePrice:    dataCore.ServicePrice,
		AverageRating:   dataCore.AverageRating,
		ServiceImageUrl: dataCore.ServiceImageUrl,
		City:            dataCore.City,
		PartnerID:       dataCore.PartnerID,
	}
}

func fromCoreAdditional(dataCore service.Additional) ServiceAdditionalResponse {
	return ServiceAdditionalResponse{
		ID:              dataCore.ID,
		AdditionalName:  dataCore.AdditionalName,
		AdditionalPrice: dataCore.AdditionalPrice,
		PartnerID:       dataCore.PartnerID,
		ServiceID:       dataCore.ServiceID,
	}
}

func fromCoreList(dataCore []service.Core) []ServiceResponse {
	var dataResponse []ServiceResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreListAdditional(dataCore []service.Additional) []ServiceAdditionalResponse {
	var dataResponse []ServiceAdditionalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreAdditional(v))
	}
	return dataResponse
}
