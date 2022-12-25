package delivery

import (
	"capstone-alta1/features/review"
)

type DataResponse struct {
	ID        uint    `json:"id"`
	Review    string  `json:"review"`
	Rating    float64 `json:"rating"`
	OrderID   uint    `json:"order_id"`
	ServiceID uint    `json:"service_id"`
	ClientID  uint    `json:"client_id"`
}

func fromCore(dataCore review.Core) DataResponse {
	return DataResponse{
		ID:        dataCore.ID,
		Review:    dataCore.Review,
		Rating:    dataCore.Rating,
		OrderID:   dataCore.OrderID,
		ServiceID: dataCore.ServiceID,
		ClientID:  dataCore.ClientID,
	}
}

func fromCoreList(dataCore []review.Core) []DataResponse {
	var dataResponse []DataResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
