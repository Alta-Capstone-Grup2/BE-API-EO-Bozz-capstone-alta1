package service

import (
	"capstone-alta1/features/auth"
	user "capstone-alta1/features/user/repository"
	"capstone-alta1/middlewares"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authData auth.RepositoryInterface
	validate *validator.Validate
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
		validate: validator.New(),
	}
}

func (service *authService) Login(dataCore auth.Core) (data auth.Core, token string, clientID uint, partnerID uint, err error) {

	if errValidate := service.validate.Struct(dataCore); errValidate != nil {
		log.Error(errValidate.Error())
		return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error validate input, please check your input.")
	}

	result, errLogin := service.authData.FindUser(dataCore.Email)
	if errLogin != nil {
		log.Error(errLogin.Error())
		if strings.Contains(errLogin.Error(), "table") {
			return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error on request, please contact your administrator.")
		} else if strings.Contains(errLogin.Error(), "found") {
			return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, email not found, please check email again.")
		} else {
			return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, other error, please contact your administrator.")
		}
	}

	errCheckPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(dataCore.Password))
	fmt.Println("\n\n Data Core User = ", dataCore)
	fmt.Println("\n\n Result User= ", result)
	if errCheckPass != nil {
		log.Error(errCheckPass.Error())
		return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, password didn't match, please check password again.")
	}

	// get data client / partner
	if result.Role == "Partner" {
		fmt.Println("Logged in as : Partner")
		dataPartner, errPartner := service.authData.FindPartner(result.ID)
		if errPartner != nil {
			log.Error(errLogin.Error())
			return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error on process, partner data not found. Please check your input.")
		}
		partnerID = dataPartner.ID
	} else if result.Role == "Client" {
		fmt.Println("Logged in as : Client")
		dataClient, errClient := service.authData.FindClient(result.ID)
		if errClient != nil {
			log.Error(errLogin.Error())
			return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error on process, partner data not found. Please check your input.")
		}
		clientID = dataClient.ID
	} else if result.Role == "Admin" {
		fmt.Println("Logged in as : Admin")
	} else {
		fmt.Println("Role not found. Role = ", result.Role)
		return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error on process get role.")
	}

	token, errToken := middlewares.CreateToken(int(result.ID), result.Name, result.Role, int(clientID), int(partnerID))
	if errToken != nil {
		log.Error(errToken.Error())
		return auth.Core{}, token, clientID, partnerID, errors.New("Failed to login, error on generate token, please try again.")
	}

	return result, token, clientID, partnerID, nil
}

func (service *authService) LoginOauth(auths auth.Oauth) (string, user.User, error) {
	token, data, err := service.authData.LoginOauth(auths)
	return token, data, err
}
