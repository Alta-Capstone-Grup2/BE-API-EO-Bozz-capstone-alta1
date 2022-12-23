package service

import (
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

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.UploadProfile(c)
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		input.CompanyImageUrl = res
	} else {
		input.CompanyImageUrl = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
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

func (service *partnerService) GetById(id int) (data partner.Core, err error) {
	data, err = service.partnerRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return partner.Core{}, err
	}

	return data, err

}

func (service *partnerService) Update(input partner.Core, id int, c echo.Context) error {
	err := service.partnerRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *partnerService) Delete(id int) error {
	// proses
	err := service.partnerRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
