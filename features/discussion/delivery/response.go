package delivery

import (
	"capstone-alta1/features/discussion"
	"time"
)

type DataResponse struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	PartnerID uint      `json:"partner_id"`
	ServiceID uint      `json:"service_id"`
	ClientID  uint      `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
}

func fromCore(dataCore discussion.Core) DataResponse {
	return DataResponse{
		ID:        dataCore.ID,
		Comment:   dataCore.Comment,
		PartnerID: dataCore.PartnerID,
		ServiceID: dataCore.ServiceID,
		ClientID:  dataCore.ClientID,
		CreatedAt: dataCore.CreatedAt,
	}
}

func fromCoreList(dataCore []discussion.Core) []DataResponse {
	var dataResponse []DataResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
