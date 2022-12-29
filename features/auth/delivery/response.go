package delivery

import (
	"capstone-alta1/features/auth"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	PartnerID uint   `json:"partner_id"`
	ClientID  uint   `json:"client_id"`
	Token     string `json:"token"`
}

func FromCore(dataCore auth.Core, token string, clientID uint, partnerID uint) UserResponse {
	return UserResponse{
		ID:        dataCore.ID,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Role:      dataCore.Role,
		ClientID:  clientID,
		PartnerID: partnerID,
		Token:     token,
	}
}
