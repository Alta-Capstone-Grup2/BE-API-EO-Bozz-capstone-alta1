package delivery

import "capstone-alta1/features/auth"

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	ClientID  uint   `json:"client_id"`
	PartnerID uint   `json:"partner_id"`
	Token     string `json:"token"`
}

func FromCore(dataCore auth.Core, token string) UserResponse {
	if dataCore.Role == "Partner" {
		ReturnPartner(dataCore, token)
	} else if dataCore.Role == "Client" {
		ReturnClient(dataCore, token)
	} else if dataCore.Role == "Admin" {
		ReturnAdmin(dataCore, token)
	}
	return UserResponse{}
}

func ReturnPartner(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Role:      dataCore.Role,
		PartnerID: dataCore.PartnerID,
		Token:     token,
	}
}

func ReturnClient(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Role:     dataCore.Role,
		ClientID: dataCore.ClientID,
		Token:    token,
	}
}

func ReturnAdmin(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
		Role:  dataCore.Role,
		Token: token,
	}
}
