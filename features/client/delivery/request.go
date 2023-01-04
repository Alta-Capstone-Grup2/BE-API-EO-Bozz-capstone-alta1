package delivery

import (
	"capstone-alta1/features/client"
)

type ClientRequest struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Gender          string `json:"gender" form:"gender"`
	Address         string `json:"address" form:"address"`
	City            string `json:"city" form:"city"`
	Phone           string `json:"phone" form:"phone"`
	ClientImageFile string `json:"client_image_file" form:"client_image_file"`
}

func toCore(input ClientRequest) client.Core {
	clientCoredata := client.Core{
		User: client.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		},
		Gender:          input.Gender,
		Address:         input.Address,
		City:            input.City,
		Phone:           input.Phone,
		ClientImageFile: input.ClientImageFile,
	}
	return clientCoredata
}

func toOrderStatus(inputComplete string, clientId uint) client.Order {
	coreInput := client.Order{
		OrderStatus: inputComplete,
		ClientID:    clientId,
	}
	return coreInput
}
