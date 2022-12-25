package delivery

import (
	"capstone-alta1/features/additional"
)

type AdditionalRequest struct {
	NameAdditional  string `json:"name_additional" form:"name_additional"`
	PriceAdditional int    `json:"price_additional" form:"price_additional"`
}

func toCore(input AdditionalRequest, id uint) additional.Core {
	clientCoredata := additional.Core{
		NameAdditional:  input.NameAdditional,
		PriceAdditional: input.PriceAdditional,
	}
	return clientCoredata
}
