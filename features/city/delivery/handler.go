package delivery

import (
	"capstone-alta1/features/city"
	"capstone-alta1/utils/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type cityDelivery struct {
	cityService city.ServiceInterface
}

func New(service city.ServiceInterface, e *echo.Echo) {
	handler := &cityDelivery{
		cityService: service,
	}

	e.GET("/city", handler.GetAll)

}

func (delivery *cityDelivery) GetAll(c echo.Context) error {
	results, err := delivery.cityService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}
