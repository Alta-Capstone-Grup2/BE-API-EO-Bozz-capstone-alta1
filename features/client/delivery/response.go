package delivery

import (
	"capstone-alta1/features/client"
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

func fromCoreList(dataCore []client.Core) []ClientResponse {
	var dataResponse []ClientResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
