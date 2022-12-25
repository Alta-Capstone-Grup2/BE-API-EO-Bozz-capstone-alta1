package delivery

import (
	"capstone-alta1/features/client"
	"time"
)

type ClientResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Phone          string `json:"phone"`
	ClientImageUrl string `json:"client_image_url"`
	UserID         uint   `json:"user_id"`
}

type ClientOrderResponse struct {
	ID            uint
	EventName     string
	StartDate     time.Time
	EndDate       time.Time
	EventLocation string
	ServiceName   string
	GrossAmmount  uint
	OrderStatus   string
	ServiceID     uint
	ClientID      uint
}

func fromCore(dataCore client.Core) ClientResponse {
	return ClientResponse{
		ID:             dataCore.ID,
		Name:           dataCore.User.Name,
		Email:          dataCore.User.Email,
		Role:           dataCore.User.Role,
		Gender:         dataCore.Gender,
		Address:        dataCore.Address,
		City:           dataCore.City,
		Phone:          dataCore.Phone,
		ClientImageUrl: dataCore.ClientImageUrl,
		UserID:         dataCore.User.ID,
	}
}

func fromCoreOrder(dataCore client.OrderCore) ClientOrderResponse {
	return ClientOrderResponse{
		ID:            dataCore.ID,
		EventName:     dataCore.EventName,
		StartDate:     dataCore.StartDate,
		EndDate:       dataCore.EndDate,
		EventLocation: dataCore.EventLocation,
		ServiceName:   dataCore.ServiceName,
		GrossAmmount:  dataCore.GrossAmmount,
		OrderStatus:   dataCore.OrderStatus,
		ServiceID:     dataCore.ServiceID,
		ClientID:      dataCore.ClientID,
	}
}

func fromCoreList(dataCore []client.Core) []ClientResponse {
	var dataResponse []ClientResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreListOrder(dataCore []client.OrderCore) []ClientOrderResponse {
	var dataResponse []ClientOrderResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreOrder(v))
	}
	return dataResponse
}
