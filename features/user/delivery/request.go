package delivery

import (
	"capstone-alta1/features/user"
)

type InsertRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func toCore(i interface{}) user.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.Core{
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.Core{
			ID:       cnv.ID,
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
		}
	case UpdatePasswordRequest:
		cnv := i.(UpdatePasswordRequest)
		return user.Core{
			Password: cnv.NewPassword,
		}
	}

	return user.Core{}
}
