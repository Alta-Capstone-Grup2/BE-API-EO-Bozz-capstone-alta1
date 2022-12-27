package delivery

import (
	"capstone-alta1/features/service"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type serviceDelivery struct {
	serviceService service.ServiceInterface
}

func New(service service.ServiceInterface, e *echo.Echo) {
	handler := &serviceDelivery{
		serviceService: service,
	}

	e.GET("/services", handler.GetAll)
	e.GET("/services/:id", handler.GetById)
	e.POST("/services", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/services/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/services/:id", handler.Delete, middlewares.JWTMiddleware())

}

func (delivery *serviceDelivery) GetAll(c echo.Context) error {
	queryName := c.QueryParam("service_name")
	queryCategory := c.QueryParam("service_category")
	queryCity := c.QueryParam("city")
	queryPrice := c.QueryParam("service_price")

	helper.LogDebug("\n\n\nULALA")

	// debug cek query param masuk
	helper.LogDebug("\n isi queryName = ", queryName)
	helper.LogDebug("\n isi queryCategory= ", queryCategory)
	helper.LogDebug("\n isi queryCity = ", queryCity)
	helper.LogDebug("\n isi queryPrice = ", queryPrice)

	results, err := delivery.serviceService.GetAll(queryName, queryCategory, queryCity, queryPrice)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *serviceDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.serviceService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *serviceDelivery) Create(c echo.Context) error {
	serviceInput := ServiceRequest{}
	errBind := c.Bind(&serviceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	ServiceImageUrl, _ := c.FormFile("service_image_file")
	if ServiceImageUrl != nil {
		urlFile, err := thirdparty.Upload(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlFile)
		serviceInput.ServiceImageUrl = urlFile
	} else {
		serviceInput.ServiceImageUrl = ""
	}

	InputPartnerID := middlewares.ExtractTokenPartnerID(c)
	dataCore := toCore(serviceInput, uint(InputPartnerID))
	err := delivery.serviceService.Create(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *serviceDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	serviceInput := ServiceRequest{}
	errBind := c.Bind(&serviceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	InputPartnerID := middlewares.ExtractTokenPartnerID(c)
	dataCore := toCore(serviceInput, uint(InputPartnerID))
	err := delivery.serviceService.Update(dataCore, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *serviceDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	var services service.Core
	if services.Order != nil {
		return c.JSON(http.StatusConflict, helper.FailedResponse("this service currently have order"))
	}
	err := delivery.serviceService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}