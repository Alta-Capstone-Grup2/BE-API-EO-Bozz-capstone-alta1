package delivery

import (
	"capstone-alta1/features/partner"
)

type PartnerResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Role               string `json:"role"`
	PICPosition        string `json:"pic_position"`
	PICPhone           string `json:"pic_phone"`
	PICAddress         string `json:"pic_address"`
	CompanyName        string `json:"company_name"`
	CompanyPhone       string `json:"company_phone"`
	CompanyCity        string `json:"company_city"`
	CompanyImageFile   string `json:"company_image_file"`
	CompanyAddress     string `json:"company_address"`
	LinkWebsite        string `json:"link_website"`
	NIBNumber          string `json:"nib_number"`
	NIBImageFile       string `json:"nib_image_file"`
	SIUPNumber         string `json:"siup_number"`
	SIUPImageFile      string `json:"siup_image_file"`
	Event1Name         string `json:"event1_name"`
	Event1ImageFile    string `json:"event1_image_file"`
	Event2Name         string `json:"event2_name"`
	Event2ImageFile    string `json:"event2_image_file"`
	Event3Name         string `json:"event3_name"`
	Event3ImageFile    string `json:"event4_image_file"`
	BankName           string `json:"bank_name"`
	BankAccountNumber  string `json:"bank_account_number"`
	BankAccountName    string `json:"bank_account_name"`
	VerificationStatus string `json:"verification_status"`
	VerificationLog    string `json:"verification_log"`
	UserID             uint   `json:"user_id"`
}

func fromCore(dataCore partner.Core) PartnerResponse {
	return PartnerResponse{
		ID:                 dataCore.ID,
		Name:               dataCore.User.Name,
		Email:              dataCore.User.Email,
		Role:               dataCore.User.Role,
		PICPosition:        dataCore.PICPosition,
		PICPhone:           dataCore.PICPhone,
		PICAddress:         dataCore.PICAddress,
		CompanyName:        dataCore.CompanyName,
		CompanyPhone:       dataCore.CompanyPhone,
		CompanyCity:        dataCore.CompanyCity,
		CompanyImageFile:   dataCore.CompanyImageFile,
		CompanyAddress:     dataCore.CompanyAddress,
		LinkWebsite:        dataCore.LinkWebsite,
		NIBNumber:          dataCore.NIBNumber,
		NIBImageFile:       dataCore.NIBImageFile,
		SIUPNumber:         dataCore.SIUPNumber,
		SIUPImageFile:      dataCore.SIUPImageFile,
		Event1Name:         dataCore.Event1Name,
		Event1ImageFile:    dataCore.Event1ImageFile,
		Event2Name:         dataCore.Event2Name,
		Event2ImageFile:    dataCore.Event2ImageFile,
		Event3Name:         dataCore.Event3Name,
		Event3ImageFile:    dataCore.Event3ImageFile,
		BankName:           dataCore.BankName,
		BankAccountNumber:  dataCore.BankAccountNumber,
		BankAccountName:    dataCore.BankName,
		VerificationStatus: dataCore.VerificationStatus,
		VerificationLog:    dataCore.VerificationLog,
		UserID:             dataCore.UserID,
	}
}

func fromCoreList(dataCore []partner.Core) []PartnerResponse {
	var dataResponse []PartnerResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
