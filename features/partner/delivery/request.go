package delivery

import (
	"capstone-alta1/features/partner"
)

type PartnerRequest struct {
	Name              string `json:"name" form:"name"`
	Email             string `json:"email" form:"email"`
	Password          string `json:"password" form:"password"`
	PICPosition       string `json:"pic_position" form:"pic_position"`
	PICPhone          string `json:"pic_phone" form:"pic_phone"`
	PICAddress        string `json:"pic_address" form:"pic_address"`
	CompanyName       string `json:"company_name" form:"company_name"`
	CompanyPhone      string `json:"company_phone" form:"company_phone"`
	CompanyCity       string `json:"company_city" form:"company_city"`
	CompanyImageFile  string `json:"company_image_file" form:"company_image_file"`
	CompanyAddress    string `json:"company_address" form:"company_address"`
	LinkWebsite       string `json:"link_website" form:"link_website"`
	NIBNumber         string `json:"nib_number" form:"nib_number"`
	NIBImageFile      string `json:"nib_image_file" form:"nib_image_file"`
	SIUPNumber        string `json:"siup_number" form:"siup_number"`
	SIUPImageFile     string `json:"siup_image_file" form:"siup_image_file"`
	Event1Name        string `json:"event1_name" form:"event1_name"`
	Event1ImageFile   string `json:"event1_image_file" form:"event1_image_file"`
	Event2Name        string `json:"event2_name" form:"event2_name"`
	Event2ImageFile   string `json:"event2_image_file" form:"event2_image_file"`
	Event3Name        string `json:"event3_name" form:"event3_name"`
	Event3ImageFile   string `json:"event3_image_file" form:"event3_image_file"`
	BankName          string `json:"bank_name" form:"bank_name"`
	BankAccountNumber string `json:"bank_account_number" form:"bank_account_number"`
	BankAccountName   string `json:"bank_account_name" form:"bank_account_name"`
}

type VerifyPartnerRequest struct {
	VerifyLog string `json:"verify_log" form:"verify_log"`
	Status    string `json:"status" form:"status"`
	PartnerID uint   `json:"partner_id" form:"partner_id"`
}

// type ConfirmOrderRequest struct {
// 	OrderStatus string `json:"order_status" form:"order_status"`
// }

func toCore(input PartnerRequest) partner.Core {
	partnerCoredata := partner.Core{
		PICPosition:       input.PICPosition,
		PICPhone:          input.PICPhone,
		PICAddress:        input.PICAddress,
		CompanyName:       input.CompanyName,
		CompanyPhone:      input.CompanyPhone,
		CompanyCity:       input.CompanyCity,
		CompanyImageFile:  input.CompanyImageFile,
		CompanyAddress:    input.CompanyAddress,
		LinkWebsite:       input.LinkWebsite,
		NIBNumber:         input.NIBNumber,
		NIBImageFile:      input.NIBImageFile,
		SIUPNumber:        input.SIUPNumber,
		SIUPImageFile:     input.SIUPImageFile,
		Event1Name:        input.Event1Name,
		Event1ImageFile:   input.Event1ImageFile,
		Event2Name:        input.Event2Name,
		Event2ImageFile:   input.Event2ImageFile,
		Event3Name:        input.Event3Name,
		Event3ImageFile:   input.Event3ImageFile,
		BankName:          input.BankName,
		BankAccountNumber: input.BankAccountNumber,
		BankAccountName:   input.BankAccountName,
		User: partner.UserCore{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		},
	}
	return partnerCoredata
}
