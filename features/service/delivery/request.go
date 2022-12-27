package delivery

import (
	"capstone-alta1/features/service"
)

type ServiceRequest struct {
	ServiceName        string  `json:"service_name" form:"service_name"`
	ServiceDescription string  `json:"service_description" form:"service_description"`
	ServiceCategory    string  `json:"service_category" form:"service_category"`
	ServicePrice       uint    `json:"service_price" form:"service_price"`
	AverageRating      float64 `json:"average_rating" form:"average_rating"`
	ServiceImageUrl    string  `json:"service_image_file" form:"service_image_file"`
	City               string  `json:"city" form:"city"`
}

func toCore(input ServiceRequest, InputPartnerID uint) service.Core {
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceDescription: input.ServiceDescription,
		ServiceCategory:    input.ServiceCategory,
		ServicePrice:       input.ServicePrice,
		AverageRating:      input.AverageRating,
		ServiceImageUrl:    input.ServiceImageUrl,
		City:               input.City,
		PartnerID:          InputPartnerID,
	}
	return coreInput
}
