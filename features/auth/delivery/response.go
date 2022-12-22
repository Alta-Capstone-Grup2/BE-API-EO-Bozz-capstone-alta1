package delivery

import "capstone-alta1/features/auth"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FromCore(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
		Role:  dataCore.Role,
		Token: token,
	}
}
