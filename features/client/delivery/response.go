package delivery

import (
	"capstone-alta1/features/client"
	"capstone-alta1/features/order"
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
}

type ClientOrderResponse struct {
	ID            uint
	EventName     string
	StartDate     time.Time
	EndDate       time.Time
	EventLocation string
	ServiceName   string
	GrossAmmount  int
	OrderStatus   string
	ServiceID     uint
	UserID        uint
}

func fromCore(dataCore client.Core) ClientResponse {
	return ClientResponse{
		ID:             dataCore.User.ID,
		Name:           dataCore.User.Name,
		Email:          dataCore.User.Email,
		Role:           dataCore.User.Role,
		Gender:         dataCore.Gender,
		Address:        dataCore.Address,
		City:           dataCore.City,
		Phone:          dataCore.Phone,
		ClientImageUrl: dataCore.ClientImageUrl,
	}
}

func fromCoreOrder(dataCore order.Core) ClientOrderResponse {
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
		UserID:        dataCore.UserID,
	}
}

func fromCoreList(dataCore []client.Core) []ClientResponse {
	var dataResponse []ClientResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreListOrder(dataCore []order.Core) []ClientOrderResponse {
	var dataResponse []ClientOrderResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreOrder(v))
	}
	return dataResponse
}
