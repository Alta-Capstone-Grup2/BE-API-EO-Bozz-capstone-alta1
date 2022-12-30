package delivery

import (
	"capstone-alta1/features/service"
	"strconv"
)

type ServiceRequest struct {
	ServiceName        string `json:"service_name" form:"service_name"`
	ServiceInclude     string `json:"service_include" form:"service_include"`
	ServiceDescription string `json:"service_description" form:"service_description"`
	ServiceCategory    string `json:"service_category" form:"service_category"`
	ServicePrice       string `json:"service_price" form:"service_price"`
	City               string `json:"city" form:"city"`
}

type ServiceUpdateRequest struct {
	ServiceName        string `json:"service_name" form:"service_name"`
	ServiceDescription string `json:"service_description" form:"service_description"`
	ServiceCategory    string `json:"service_category" form:"service_category"`
	ServicePrice       string `json:"service_price" form:"service_price"`
	AverageRating      string `json:"average_rating" form:"average_rating"`
	City               string `json:"city" form:"city"`
}

type ServiceAdditionalRequest struct {
	ServiceID         uint `json:"service_id" form:"service_id"`
	AdditionalRequest []AdditionalRequest
}
type AdditionalRequest struct {
	AdditionalID uint `json:"additional_id" form:"additional_id"`
}

func toCore(input ServiceRequest, InputPartnerID uint) service.Core {
	priceInt, _ := strconv.Atoi(input.ServicePrice)
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceInclude:     input.ServiceInclude,
		ServiceDescription: input.ServiceDescription,
		ServiceCategory:    input.ServiceCategory,
		ServicePrice:       uint(priceInt),
		City:               input.City,
		PartnerID:          InputPartnerID,
	}
	return coreInput
}

func toCoreUpdate(input ServiceUpdateRequest, InputPartnerID uint) service.Core {
	priceInt, _ := strconv.Atoi(input.ServicePrice)
	ratingInt, _ := strconv.Atoi(input.AverageRating)
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceDescription: input.ServiceDescription,
		ServiceCategory:    input.ServiceCategory,
		ServicePrice:       uint(priceInt),
		AverageRating:      float64(ratingInt),
		City:               input.City,
		PartnerID:          InputPartnerID,
	}
	return coreInput
}

func toCoreAdditional(input ServiceAdditionalRequest) service.ServiceAdditional {
	inputAdditional := AdditionalRequest{}
	coreInput := service.ServiceAdditional{
		ServiceID:    input.ServiceID,
		AdditionalID: inputAdditional.AdditionalID,
	}
	return coreInput
}
