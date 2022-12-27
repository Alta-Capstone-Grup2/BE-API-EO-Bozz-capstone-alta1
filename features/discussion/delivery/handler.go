package delivery

import (
	"capstone-alta1/features/discussion"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DiscussionDelivery struct {
	discussionService discussion.ServiceInterface
}

func New(service discussion.ServiceInterface, e *echo.Echo) {
	handler := &DiscussionDelivery{
		discussionService: service,
	}

	e.GET("/discussions", handler.GetAll)
	e.GET("/discussions/:id", handler.GetById)
	e.POST("/discussions", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/discussions/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/discussions/:id", handler.Delete, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *DiscussionDelivery) GetAll(c echo.Context) error {
	results, err := delivery.discussionService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *DiscussionDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.discussionService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *DiscussionDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	fmt.Println("Cek discussion data = ", userInput)
	dataCore := toCore(userInput)

	fmt.Println("Cek discussion data = ", dataCore)
	// paratner id client id dr token // di comment dlu
	// dataCore.PartnerID = uint(middlewares.ExtractTokenPartnerID(c))
	// dataCore.ClientID = uint(middlewares.ExtractTokenClientID(c))

	fmt.Println("Cek discussion data = ", dataCore)

	err := delivery.discussionService.Create(dataCore, c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *DiscussionDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userInput := UpdateRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	// data, errGet := delivery.discussionService.GetById(id)
	// if errGet != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	// }

	// if userId != int(data.ClientID) {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	dataCore := toCore(userInput)

	// paratner id client id dr token
	dataCore.PartnerID = uint(middlewares.ExtractTokenPartnerID(c))
	dataCore.ClientID = uint(middlewares.ExtractTokenClientID(c))

	err := delivery.discussionService.Update(dataCore, uint(id), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *DiscussionDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// validasi data di proses oleh user ybs
	userId := middlewares.ExtractTokenUserId(c)
	if userId < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load user id from JWT token, please check again."))
	}
	// data, errGet := delivery.discussionService.GetById(id)
	// if errGet != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	// }

	// if userId != int(data.ClientID) {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	err := delivery.discussionService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
