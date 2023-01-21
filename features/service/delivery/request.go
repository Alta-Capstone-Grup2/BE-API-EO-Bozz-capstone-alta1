package delivery

import (
	"capstone-alta1/features/service"
	"strconv"
)

type ServiceRequest struct {
	ServiceName        string `json:"service_name" form:"service_name" `
	ServiceIncluded    string `json:"service_included" form:"service_included"`
	ServiceDescription string `json:"service_description" form:"service_description"`
	ServiceCategory    string `json:"service_category" form:"service_category"`
	ServicePrice       string `json:"service_price" form:"service_price"`
	City               string `json:"city" form:"city"`
}

type ServiceUpdateRequest struct {
	ServiceName        string `json:"service_name" form:"service_name"`
	ServiceIncluded    string `json:"service_included" form:"service_included"`
	ServiceDescription string `json:"service_description" form:"service_description"`
	ServiceCategory    string `json:"service_category" form:"service_category"`
	ServicePrice       string `json:"service_price" form:"service_price"`
	City               string `json:"city" form:"city"`
}

type ServiceAdditionalRequest struct {
	ServiceID   uint                `json:"service_id" form:"service_id"`
	Additionals []AdditionalRequest `json:"additionals" form:"additionals"`
}
type AdditionalRequest struct {
	AdditionalID uint `json:"additional_id" form:"additional_id"`
}

func toCore(input ServiceRequest, InputPartnerID uint) service.Core {
	priceInt, _ := strconv.Atoi(input.ServicePrice)
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceIncluded:    input.ServiceIncluded,
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
	coreInput := service.Core{
		ServiceName:        input.ServiceName,
		ServiceIncluded:    input.ServiceIncluded,
		ServiceDescription: input.ServiceDescription,
		ServiceCategory:    input.ServiceCategory,
		ServicePrice:       uint(priceInt),
		City:               input.City,
		PartnerID:          InputPartnerID,
	}
	return coreInput
}

// func toCoreAdditional(input AdditionalRequest) service.ServiceAdditional {
// 	coreInput := service.ServiceAdditional{
// 		AdditionalID: input.AdditionalID,
// 	}
// 	return coreInput
// }

func toServiceAdditionalList(requestData ServiceAdditionalRequest) []service.ServiceAdditional {
	var dataCore []service.ServiceAdditional
	for _, v := range requestData.Additionals {
		dataCore = append(dataCore, service.ServiceAdditional{
			ServiceID:    requestData.ServiceID,
			AdditionalID: v.AdditionalID,
		})
	}
	return dataCore
}

// func toAdditionalList(requestData []AdditionalRequest) []service.ServiceAdditional {
// 	var dataCore []service.ServiceAdditional
// 	for _, v := range requestData {
// 		dataCore = append(dataCore, toCoreAdditional(v))
// 	}
// 	return dataCore
// }
