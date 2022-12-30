package delivery

import (
	"capstone-alta1/features/service"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
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
	e.GET("/services/:id/additionals", handler.GetAdditionalById)
	e.GET("/services/:id/reviews", handler.GetReviewById)
	e.GET("/services/:id/discussions", handler.GetDiscussionById)
	e.POST("/services/additionals", handler.AddAdditionalToService, middlewares.JWTMiddleware())
	e.POST("/services", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/services/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/services/:id", handler.Delete, middlewares.JWTMiddleware())
	e.POST("/services/:id/availability", handler.CheckAvailability, middlewares.JWTMiddleware())

}

func (delivery *serviceDelivery) GetAll(c echo.Context) error {
	queryName := c.QueryParam("service_name")
	queryCategory := c.QueryParam("service_category")
	queryCity := c.QueryParam("city")
	queryMinPrice := c.QueryParam("min_price")
	queryMaxPrice := c.QueryParam("min_price")

	helper.LogDebug("\n\n\nULALA")

	// debug cek query param masuk
	helper.LogDebug("\n isi queryName = ", queryName)
	helper.LogDebug("\n isi queryCategory= ", queryCategory)
	helper.LogDebug("\n isi queryCity = ", queryCity)
	helper.LogDebug("\n isi queryMinPrice = ", queryMinPrice)
	helper.LogDebug("\n isi queryMaxPrice = ", queryMaxPrice)

	results, err := delivery.serviceService.GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice)
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

	dataResponse := fromCoreGetById(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *serviceDelivery) Create(c echo.Context) error {
	serviceInput := ServiceRequest{}
	errBind := c.Bind(&serviceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	InputPartnerID := middlewares.ExtractTokenPartnerID(c)
	dataCore := toCore(serviceInput, uint(InputPartnerID))
	err := delivery.serviceService.Create(dataCore, c)
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

	serviceInput := ServiceUpdateRequest{}
	errBind := c.Bind(&serviceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	InputPartnerID := middlewares.ExtractTokenPartnerID(c)
	dataCore := toCoreUpdate(serviceInput, uint(InputPartnerID))
	err := delivery.serviceService.Update(dataCore, uint(id), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
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

func (delivery *serviceDelivery) GetAdditionalById(c echo.Context) error {
	idParam := c.Param("id")
	serviceId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.serviceService.GetAdditionalById(uint(serviceId))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreListAdditional(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *serviceDelivery) AddAdditionalToService(c echo.Context) error {
	serviceInput := ServiceAdditionalRequest{}
	errBind := c.Bind(&serviceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCoreAdditional(serviceInput)
	err := delivery.serviceService.AddAdditionalToService(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *serviceDelivery) GetReviewById(c echo.Context) error {
	idParam := c.Param("id")
	serviceId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.serviceService.GetReviewById(uint(serviceId))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreListReview(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *serviceDelivery) GetDiscussionById(c echo.Context) error {
	idParam := c.Param("id")
	serviceId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.serviceService.GetDiscussionById(uint(serviceId))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreListDiscussion(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *serviceDelivery) CheckAvailability(c echo.Context) error {
	idParam := c.Param("id")
	serviceId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	queryStart := c.QueryParam("start_date")
	queryEnd := c.QueryParam("end_date")
	helper.LogDebug("\n isi queryStart = ", queryStart)
	helper.LogDebug("\n isi queryEnd= ", queryEnd)

	data, err := delivery.serviceService.CheckAvailability(uint(serviceId), queryStart, queryEnd)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("Success create data", fromCoreAvailability(data)))
}
