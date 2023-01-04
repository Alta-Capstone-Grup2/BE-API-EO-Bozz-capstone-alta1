package delivery

import (
	"capstone-alta1/features/order"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type orderDelivery struct {
	orderService order.ServiceInterface
}

func New(order order.ServiceInterface, e *echo.Echo) {
	handler := &orderDelivery{
		orderService: order,
	}

	e.POST("/orders", handler.Create, middlewares.JWTMiddleware())
	e.GET("/orders", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/orders/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/orders/:id", handler.UpdateStatusCancel, middlewares.JWTMiddleware())
	e.PUT("/orders/:id/payout", handler.UpdateStatusPayout, middlewares.JWTMiddleware())
	e.POST("/orders/payment_notifications", handler.UpdateMidtrans())
}

func (delivery *orderDelivery) Create(c echo.Context) error {
	// authorization
	userRole := middlewares.ExtractTokenUserRole(c)
	if userRole != "Client" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Error process request. This action only for Client"))
	}

	orderInput := OrderRequest{}
	errBind := c.Bind(&orderInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		helper.LogDebug("Order - handler - Create | Error binding data. Error  = ", errBind.Error)
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. Please check again."))
	}

	//input validation
	if errValidateInput := Validate(orderInput); errValidateInput != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect input. "+errValidateInput.Error()))
	}

	//validate startdate enddate format
	if errFormatStart := helper.ValidateDateFormat(orderInput.StartDate); errFormatStart != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errFormatStart.Error()))
	}
	if errFormatEnd := helper.ValidateDateFormat(orderInput.EndDate); errFormatEnd != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(errFormatEnd.Error()))
	}

	inputClientID := middlewares.ExtractTokenClientID(c)
	dataCore := toCore(orderInput, uint(inputClientID))
	dataDetailOrder := toDetailOrderList(orderInput.OrderDetails)

	helper.LogDebug("Order - handler - Create | Data bind  = ", helper.ConvToJson(orderInput))
	helper.LogDebug("Order - handler - Create | Data core order  = ", helper.ConvToJson(dataCore))
	helper.LogDebug("Order - handler - Create | Data detail order  = ", helper.ConvToJson(dataDetailOrder))

	// result, err := delivery.orderService.Create(dataCore, dataDetailOrder)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "Error:Field validation") {
	// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
	// 	}
	// 	if strings.Contains(err.Error(), "Service Data or Additional Data Not Found. Please Check your input") {
	// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert data. "+err.Error()))
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	// }

	//dataResponse := fromCoreToPayment(result)

	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("Success create data", nil))
}

func (delivery *orderDelivery) GetAll(c echo.Context) error {
	query := c.QueryParam("event_name")
	helper.LogDebug("\n isi query = ", query)

	results, err := delivery.orderService.GetAll(query)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromViewCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *orderDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, results2, err := delivery.orderService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results, results2)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *orderDelivery) UpdateStatusCancel(c echo.Context) error {
	idParam := c.Param("id")
	orderId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	// orderInput := OrderStatusRequest{}
	// errBind := c.Bind(&orderInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	// if errBind != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	// }
	orderInput := "Cancel Order"
	dataCore := toCoreStatus(orderInput, uint(orderId))
	err := delivery.orderService.UpdateStatusCancel(dataCore, uint(orderId))
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

func (delivery *orderDelivery) UpdateStatusPayout(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	err := delivery.orderService.UpdateStatusPayout(uint(id), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

// CALLBACK MIDTRANS
func (delivery *orderDelivery) UpdateMidtrans() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateMidtransRequest
		errBind := c.Bind(&input)
		if errBind != nil {
			helper.LogDebug("Order - handler - UpdateMidtrans | Error binding data. Error  = ", errBind.Error)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data."))
		}

		helper.LogDebug("Order - handler - UpdateMidtrans | Binding data  = ", input)

		res := toUpdateMidtrans(input)
		delivery.orderService.UpdateMidtrans(res)
		return c.JSON(http.StatusAccepted, helper.SuccessResponse("Success update order data."))
	}
}
