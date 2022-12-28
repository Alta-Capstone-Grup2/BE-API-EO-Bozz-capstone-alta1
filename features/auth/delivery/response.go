package delivery

import "capstone-alta1/features/auth"

type UserAdminResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UserPartnerResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	PartnerID uint   `json:"partner_id"`
	Token     string `json:"token"`
}

type UserClientResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ClientID uint   `json:"client_id"`
	Token    string `json:"token"`
}

func FromCore(dataCore auth.Core, token string) UserAdminResponse {
	if dataCore.Role == "Partner" {
		ReturnPartner(dataCore, token)
	} else if dataCore.Role == "Client" {
		ReturnClient(dataCore, token)
	}
	return ReturnAdmin(dataCore, token)
}

func ReturnPartner(dataCore auth.Core, token string) UserPartnerResponse {
	return UserPartnerResponse{
		ID:        dataCore.ID,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Role:      dataCore.Role,
		PartnerID: dataCore.PartnerID,
		Token:     token,
	}
}

func ReturnClient(dataCore auth.Core, token string) UserClientResponse {
	return UserClientResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Role:     dataCore.Role,
		ClientID: dataCore.ClientID,
		Token:    token,
	}
}

func ReturnAdmin(dataCore auth.Core, token string) UserAdminResponse {
	return UserAdminResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
		Role:  dataCore.Role,
		Token: token,
	}
}
