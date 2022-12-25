package delivery

import (
	"capstone-alta1/features/review"
)

type InsertRequest struct {
	Review    string  `json:"review" form:"review"`
	Rating    float64 `json:"rating" form:"rating"`
	OrderID   uint    `json:"order_id" form:"order_id"`
	ServiceID uint    `json:"service_id" form:"service"`
}

type UpdateRequest struct {
	Review    string  `json:"review" form:"review"`
	Rating    float64 `json:"rating" form:"rating"`
	OrderID   uint    `json:"order_id" form:"order_id"`
	ServiceID uint    `json:"service_id" form:"service"`
}

func toCore(i interface{}, userId uint) review.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return review.Core{
			Review:    cnv.Review,
			Rating:    cnv.Rating,
			OrderID:   cnv.OrderID,
			ServiceID: cnv.ServiceID,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return review.Core{
			Review:    cnv.Review,
			Rating:    cnv.Rating,
			OrderID:   cnv.OrderID,
			ServiceID: cnv.ServiceID,
		}
	}

	return review.Core{}
}
