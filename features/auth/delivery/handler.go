package delivery

import (
	"capstone-alta1/features/auth"
	"capstone-alta1/utils/helper"
	oauth "capstone-alta1/utils/thirdparty"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthHandler{
		authService: service,
	}
	e.POST("/login", handler.Login)
	e.GET("/login/oauth/google", LoginOauthGoogle)
	e.GET("/callback/oauth/google", handler.CallbackOauthGoogle)
}

func (handler *AuthHandler) Login(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to bind data."))
	}

	dataCore := ToCore(userInput)
	result, token, clientID, partnerID, err := handler.authService.Login(dataCore)

	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to Login. "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Success.", FromCore(result, token, clientID, partnerID)))
}

func LoginOauthGoogle(c echo.Context) error {
	// var w http.ResponseWriter
	// var r *http.Request
	// Create oauthState cookie
	oauthState := oauth.GenerateStateOauthCookie(c)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	u := oauth.AuthConfig().AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, u)
	return c.JSON(http.StatusOK, "success")
}

func (handler *AuthHandler) CallbackOauthGoogle(c echo.Context) error {
	// Read oauthState from Cookie
	oauthState, err := c.Cookie("oauthstate")

	helper.LogDebug("Auth - Handler - CallbackOauthGoogle | err cookie oauthstate. Error = ", err)
	helper.LogDebug("Auth - Handler - CallbackOauthGoogle | err cookie oauthstate. Oauthstate = ", helper.ConvToJson(oauthState))
	helper.LogDebug("Auth - Handler - CallbackOauthGoogle | err cookie oauthstate. from value  = ", c.FormValue("state"))

	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return errors.New("error callback")
	}

	data, err := oauth.GetUserDataFromGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return err
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	// fmt.Fprintf(c, "UserInfo: %s\n", data)
	var google auth.Oauth
	errUnmarshal := json.Unmarshal(data, &google)
	if errUnmarshal != nil {
		log.Fatal("error unmarshal")
	}

	token, dataClient, err := handler.authService.LoginOauth(google)
	if err != nil {
		helper.LogDebug("Auth - Handler - CallbackOauthGoogle | err process login oauth Error = ", err)
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	call_back_url := fmt.Sprintf("%s%s?token=%s&name=%s&user_id=%d&client_id=%d&partner_id=%d&role=%s", os.Getenv("SERVER_FRONTEND"), os.Getenv("REDIRECT_OAUTH_LOGIN"), token, dataClient.User.Name, dataClient.User.ID, dataClient.ID, 0, dataClient.User.Role)

	helper.LogDebug("Auth - Handler - CallbackOauthGoogle | Call Back URL = ", call_back_url)

	return c.Redirect(http.StatusTemporaryRedirect, call_back_url)

}
