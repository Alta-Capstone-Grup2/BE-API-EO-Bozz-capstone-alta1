package delivery

import (
	"capstone-alta1/features/review"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReviewDelivery struct {
	reviewService review.ServiceInterface
}

func New(service review.ServiceInterface, e *echo.Echo) {
	handler := &ReviewDelivery{
		reviewService: service,
	}

	e.GET("/reviews", handler.GetAll)
	e.GET("/reviews/:id", handler.GetById)
	e.POST("/reviews", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/reviews/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/reviews/:id", handler.Delete, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *ReviewDelivery) GetAll(c echo.Context) error {
	query := c.QueryParam("title")
	helper.LogDebug("isi query = ", query)
	results, err := delivery.reviewService.GetAll(query)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *ReviewDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.reviewService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *ReviewDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole != "Client" {
		return c.JSON(http.StatusNotAcceptable, helper.FailedResponse("other than a user with a client role can't give a review"))
	}

	clientID := middlewares.ExtractTokenClientID(c)
	var review review.Core
	//validasi client have one review on one order
	if clientID == int(review.ClientID) && userInput.OrderID == review.OrderID {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("This Client has already given a review to this order, please input other order id"))
	}

	dataCore := toCore(userInput, uint(clientID))
	err := delivery.reviewService.Create(dataCore, uint(clientID), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ReviewDelivery) Update(c echo.Context) error {
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
	clientID := middlewares.ExtractTokenClientID(c)
	if clientID < 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed load client id from JWT token, please check again."))
	}
	// data, errGet := delivery.reviewService.GetById(id)
	// if errGet != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	// }

	// if userId != int(data.ClientID) {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	dataCore := toCore(userInput, uint(clientID))
	err := delivery.reviewService.Update(dataCore, uint(id), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

func (delivery *ReviewDelivery) Delete(c echo.Context) error {
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
	// data, errGet := delivery.reviewService.GetById(id)
	// if errGet != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse(errGet.Error()))
	// }

	// if userId != int(data.ClientID) {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed process data, data must be yours."))
	// }

	// process
	err := delivery.reviewService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
