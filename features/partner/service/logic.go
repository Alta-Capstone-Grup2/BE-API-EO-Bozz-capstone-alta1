package service

import (
	cfg "capstone-alta1/config"
	"capstone-alta1/features/partner"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type partnerService struct {
	partnerRepository partner.RepositoryInterface
	validate          *validator.Validate
}

func New(repo partner.RepositoryInterface) partner.ServiceInterface {
	return &partnerService{
		partnerRepository: repo,
		validate:          validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *partnerService) Create(input partner.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	input.User.Role = "Partner"
	input.VerificationStatus = "Not Verified"

	// datetime layout
	layoutDefault := "2006-01-02 15:04:05"
	//init the loc
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//set timezone,
	now := time.Now().In(loc).Format(layoutDefault)

	input.VerificationLog = now + " Partner Register."

	// validasi email harus unik
	data, errFindEmail := service.partnerRepository.FindUser(input.User.Email)
	// helper.LogDebug("\n\n\n find email input  ", input.Email)
	// helper.LogDebug("\n\n\n find email data  ", data.Email)

	if data.User.Email == input.User.Email {
		return errors.New("Email " + input.User.Email + " already exist. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	// upload file
	var errUpload error
	input.CompanyImageFile, errUpload = thirdparty.Upload(c, cfg.COMPANY_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	input.NIBImageFile, errUpload = thirdparty.Upload(c, cfg.NIB_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	input.SIUPImageFile, errUpload = thirdparty.Upload(c, cfg.SIUP_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	input.Event1ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT1_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	input.Event2ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT2_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	input.Event3ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT3_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	// Encrypt
	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.User.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return err
	}

	input.User.Password = string(bytePass)

	// Process
	errCreate := service.partnerRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return err
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *partnerService) GetAll(query string) (data []partner.Core, err error) {
	if query == "" {
		data, err = service.partnerRepository.GetAll()
	} else {
		data, err = service.partnerRepository.GetAllWithSearch(query)
	}

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *partnerService) GetById(id uint) (data partner.Core, err error) {
	data, err = service.partnerRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return partner.Core{}, err
	}

	return data, err

}

func (service *partnerService) Update(input partner.Core, partnerId, userId uint, c echo.Context) error {
	var partnerData partner.Core
	// upload file
	var errUpload error
	input.CompanyImageFile, errUpload = thirdparty.Upload(c, cfg.COMPANY_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.CompanyImageFile == "" {
		input.CompanyImageFile = partnerData.CompanyImageFile
	}
	if errUpload != nil {
		return errUpload
	}
	input.NIBImageFile, errUpload = thirdparty.Upload(c, cfg.NIB_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.NIBImageFile == "" {
		input.NIBImageFile = partnerData.NIBImageFile
	}
	if errUpload != nil {
		return errUpload
	}
	input.SIUPImageFile, errUpload = thirdparty.Upload(c, cfg.SIUP_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.SIUPImageFile == "" {
		input.SIUPImageFile = partnerData.SIUPImageFile
	}
	if errUpload != nil {
		return errUpload
	}
	input.Event1ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT1_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.Event1ImageFile == "" {
		input.Event1ImageFile = partnerData.Event1ImageFile
	}
	if errUpload != nil {
		return errUpload
	}
	input.Event2ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT2_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.Event2ImageFile == "" {
		input.Event2ImageFile = partnerData.Event2ImageFile
	}
	if errUpload != nil {
		return errUpload
	}
	input.Event3ImageFile, errUpload = thirdparty.Upload(c, cfg.EVENT3_IMAGE_FILE, cfg.PARTNER_FOLDER)
	if input.Event3ImageFile == "" {
		input.Event3ImageFile = partnerData.Event3ImageFile
	}
	if errUpload != nil {
		return errUpload
	}

	err := service.partnerRepository.Update(input, partnerId, userId)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *partnerService) Delete(partnerId, userId uint) error {
	// proses
	err := service.partnerRepository.Delete(partnerId, userId)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

func (service *partnerService) GetServices(partnerID uint) (data []partner.ServiceCore, err error) {

	data, err = service.partnerRepository.GetServices(partnerID)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *partnerService) GetOrders(partnerID uint) (data []partner.OrderCore, err error) {

	data, err = service.partnerRepository.GetOrders(partnerID)

	helper.LogDebug("Partner - logic - get orders | partner id = ", partnerID)
	helper.LogDebug("Partner - logic - get orders |  data = ", data)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *partnerService) GetAdditionals(partnerID uint) (data []partner.AdditionalCore, err error) {

	data, err = service.partnerRepository.GetAdditionals(partnerID)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}
func (service *partnerService) GetPartnerRegisterData(queryCompanyName, queryPICName, queryPartnerStatus string) (data []partner.Core, err error) {

	data, err = service.partnerRepository.GetPartnerRegisterData(queryCompanyName, queryPICName, queryPartnerStatus)

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	// helper.LogDebug("Partnre Service - GetPartnerRegisterData", data)

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}
func (service *partnerService) GetPartnerRegisterDataByID(partnerID uint) (data partner.Core, err error) {
	data, err = service.partnerRepository.GetPartnerRegisterDataByID(partnerID)
	if err != nil {
		helper.LogDebug(err.Error())
		return partner.Core{}, err
	}

	return data, err
}
func (service *partnerService) UpdatePartnerVerifyStatus(verificationLog, verificationStatus string, partnerID uint) (err error) {
	// datetime layout
	layoutDefault := "2006-01-02 15:04:05"
	//init the loc
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//set timezone,
	now := time.Now().In(loc).Format(layoutDefault)

	verificationLog = now + " " + verificationLog

	// validasi status tidak bisa diubah jika sudah confirmed
	dataResult, err := service.partnerRepository.GetById(partnerID)
	if err != nil {
		helper.LogDebug(err.Error())
		return err
	}

	if verificationStatus == cfg.PARTNER_VERIFICATION_STATUS_VERIFIED && dataResult.VerificationStatus == cfg.PARTNER_VERIFICATION_STATUS_VERIFIED {
		return errors.New("Failed. Cannot update status that already Verified.")
	}

	// proses
	err = service.partnerRepository.UpdatePartnerVerifyStatus(verificationLog, verificationStatus, partnerID)
	if err != nil {
		helper.LogDebug(err.Error())
		return err
	}

	return err
}
func (service *partnerService) UpdateOrderConfirmStatus(orderID uint, partnerID uint) (err error) {
	err = service.partnerRepository.UpdateOrderConfirmStatus(orderID, partnerID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return err
}
