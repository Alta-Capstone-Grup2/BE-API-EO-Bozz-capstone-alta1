package delivery

import (
	"capstone-alta1/features/client"
	user "capstone-alta1/features/user"
)

type ClientRequest struct {
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	Gender         string `json:"gender" form:"gender"`
	Address        string `json:"address" form:"address"`
	City           string `json:"city" form:"city"`
	Phone          string `json:"phone" form:"phone"`
	ClientImageUrl string `json:"client_image_url" form:"client_image_url"`
}

func toCore(input ClientRequest) client.Core {
	clientCoredata := client.Core{
		User: user.Core{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		},
		Gender:         input.Gender,
		Address:        input.Address,
		City:           input.City,
		Phone:          input.Phone,
		ClientImageUrl: input.ClientImageUrl,
	}
	return clientCoredata
}
