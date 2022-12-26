package delivery

import (
	"capstone-alta1/features/discussion"
)

type InsertRequest struct {
	Comment   string `json:"comment" form:"comment"`
	PartnerID uint   `json:"partner_id" form:"partner_id"`
	ClientID  uint   `json:"client_id" form:"client_id"`
	ServiceID uint   `json:"service_id" form:"service_id"`
}

type UpdateRequest struct {
	Comment   string `json:"comment" form:"comment"`
	PartnerID uint   `json:"partner_id" form:"partner_id"`
	ClientID  uint   `json:"client_id" form:"client_id"`
	ServiceID uint   `json:"service_id" form:"service_id"`
}

func toCore(i interface{}) discussion.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return discussion.Core{
			Comment:   cnv.Comment,
			PartnerID: cnv.PartnerID,
			ClientID:  cnv.ClientID,
			ServiceID: cnv.ServiceID,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return discussion.Core{
			Comment:   cnv.Comment,
			PartnerID: cnv.PartnerID,
			ClientID:  cnv.ClientID,
			ServiceID: cnv.ServiceID,
		}
	}

	return discussion.Core{}
}
