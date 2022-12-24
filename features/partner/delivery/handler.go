package delivery

import (
	"capstone-alta1/features/partner"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type PartnerDelivery struct {
	partnerService partner.ServiceInterface
}

func New(service partner.ServiceInterface, e *echo.Echo) {
	handler := &PartnerDelivery{
		partnerService: service,
	}

	e.GET("/partners", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/partners/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/partners", handler.Create)
	e.PUT("/partners", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/partners", handler.Delete, middlewares.JWTMiddleware())

}

func (delivery *PartnerDelivery) GetAll(c echo.Context) error {
	query := c.QueryParam("name")
	helper.LogDebug("isi query = ", query)
	results, err := delivery.partnerService.GetAll(query)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *PartnerDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.partnerService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *PartnerDelivery) Create(c echo.Context) error {
	userInput := PartnerRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.partnerService.Create(dataCore, c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot empty. Details : "+err.Error()))
		}
		if strings.Contains(err.Error(), "Please pick another email.") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert data "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *PartnerDelivery) Update(c echo.Context) error {
	idUser := middlewares.ExtractTokenUserId(c)
	userInput := PartnerRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.partnerService.Update(dataCore, idUser, c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *PartnerDelivery) Delete(c echo.Context) error {
	idUser := middlewares.ExtractTokenUserId(c)
	err := delivery.partnerService.Delete(idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
