package delivery

import (
	"capstone-alta1/features/additional"
)

type AdditionalRequest struct {
	AdditionalName  string `json:"additional_name" form:"additional_name"`
	AdditionalPrice uint   `json:"additional_price" form:"additional_price"`
}

func toCore(input AdditionalRequest, partnerID uint) additional.Core {
	clientCoredata := additional.Core{
		AdditionalName:  input.AdditionalName,
		AdditionalPrice: input.AdditionalPrice,
		PartnerID:       partnerID,
	}
	return clientCoredata
}
