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
	// ServiceImageFile   string  `json:"service_image_file" form:"service_image_file"`
	City string `json:"city" form:"city"`
}

type ServiceAdditionalRequest struct {
	ServiceID    uint `json:"service_id" form:"service_id"`
	AdditionalID uint `json:"additional_id" form:"additional_id"`
}

func toCore(input ServiceRequest, InputPartnerID uint) service.Core {
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceDescription: input.ServiceDescription,
		ServiceCategory:    input.ServiceCategory,
		ServicePrice:       input.ServicePrice,
		AverageRating:      input.AverageRating,
		// ServiceImageFile:   input.ServiceImageFile,
		City:      input.City,
		PartnerID: InputPartnerID,
	}
	return coreInput
}

func toCoreAdditional(input ServiceAdditionalRequest) service.ServiceAdditional {
	coreInput := service.ServiceAdditional{
		ServiceID:    input.ServiceID,
		AdditionalID: input.AdditionalID,
	}
	return coreInput
}
