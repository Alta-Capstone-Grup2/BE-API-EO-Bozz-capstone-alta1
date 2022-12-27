package delivery

import "capstone-alta1/features/service"

type ServiceResponse struct {
	ID              uint    `json:"id"`
	ServiceName     string  `json:"service_name"`
	ServiceCategory string  `json:"service_category"`
	ServicePrice    uint    `json:"service_proce"`
	AverageRating   float64 `json:"average_rating"`
	ServiceImageUrl string  `json:"service_image_file"`
	City            string  `json:"city"`
	PartnerID       uint    `json:"partner_id"`
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

func fromCoreList(dataCore []service.Core) []ServiceResponse {
	var dataResponse []ServiceResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
