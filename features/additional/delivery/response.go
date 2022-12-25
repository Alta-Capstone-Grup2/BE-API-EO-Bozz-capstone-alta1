package delivery

import (
	"capstone-alta1/features/additional"
)

type AdditionalResponse struct {
	ID              uint   `json:"id"`
	NameAdditional  string `json:"name_additional"`
	PriceAdditional int    `json:"price_additional"`
	PartnerID       uint   `json:"partner_id"`
}

func fromCore(dataCore additional.Core) AdditionalResponse {
	return AdditionalResponse{
		ID:              dataCore.ID,
		NameAdditional:  dataCore.NameAdditional,
		PriceAdditional: dataCore.PriceAdditional,
		PartnerID:       dataCore.PartnerID,
	}
}

func fromCoreList(dataCore []additional.Core) []AdditionalResponse {
	var dataResponse []AdditionalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
