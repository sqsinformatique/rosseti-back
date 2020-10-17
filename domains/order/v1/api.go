package orderv1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o OrderV1) OrderPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var order models.Order
	err = ec.Bind(&order)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE USER FAILED %+v", &order)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	orderData, err := o.CreateOrder(&order)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE ORDER FAILED %+v", &order)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.CreateFailed(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrderGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	orderID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	orderData, err := o.GetOrderByID(orderID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrderSignSuperviserPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	orderID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	orderData, err := o.GetOrderByID(orderID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	if orderData.SuperviserSign != "" {
		err = errors.Errorf("already signed by supervisor")
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusForbidden,
			httpsrv.Forbidden(err),
		)
	}

	dataForSign := make(map[string]interface{})
	dataForSign["staff"] = orderData.StaffID
	dataForSign["supervisor"] = orderData.SuperviserID
	dataForSign["tech_tasks"] = orderData.TechTasks
	dataForSign["created_at"] = orderData.CreatedAt

	orderData.SuperviserSign, err = o.profileV1.SignDataByID(int64(orderData.SuperviserID), &dataForSign)
	if err != nil {
		hndlLog.Err(err).Msgf("FAILED SIGN BY Supervisor, id %d", orderID)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}
	orderData.SuperviserSignEt.Time = time.Now()
	orderData.SuperviserSignEt.Valid = true

	jsonObject, err := json.Marshal(orderData)
	if err != nil {
		hndlLog.Err(err).Msgf("marshal, id %d", orderID)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	orderData, err = o.UpdateOrderByID(orderID, &jsonObject)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", orderID, string(jsonObject))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrderSignStaffPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	orderID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	orderData, err := o.GetOrderByID(orderID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	if orderData.SuperviserSign == "" {
		err = errors.Errorf("not signed by supervisor")
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusForbidden,
			httpsrv.Forbidden(err),
		)
	}

	if orderData.StaffSign != "" {
		err = errors.Errorf("already signed by staff")
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", orderID)

		return ec.JSON(
			http.StatusForbidden,
			httpsrv.Forbidden(err),
		)
	}

	dataForSign := make(map[string]interface{})
	dataForSign["staff"] = orderData.StaffID
	dataForSign["supervisor"] = orderData.SuperviserID
	dataForSign["tech_tasks"] = orderData.TechTasks
	dataForSign["created_at"] = orderData.CreatedAt

	orderData.StaffSign, err = o.profileV1.SignDataByID(int64(orderData.StaffID), &dataForSign)
	if err != nil {
		hndlLog.Err(err).Msgf("FAILED SIGN BY Supervisor, id %d", orderID)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	orderData.StaffSignEt.Time = time.Now()
	orderData.StaffSignEt.Valid = true

	jsonObject, err := json.Marshal(orderData)
	if err != nil {
		hndlLog.Err(err).Msgf("marshal, id %d", orderID)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	orderData, err = o.UpdateOrderByID(orderID, &jsonObject)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", orderID, string(jsonObject))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrdersGetByUserIDHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	userID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	orderData, err := o.GetOrdersByUserID(userID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", userID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrderPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	orderID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	var bodyBytes []byte
	if ec.Request().Body != nil {
		bodyBytes, err = ioutil.ReadAll(ec.Request().Body)

		ec.Request().Body.Close()

		if err != nil {
			hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %d", orderID)

			return ec.JSON(
				http.StatusBadRequest,
				httpsrv.BadRequest(err),
			)
		}
	}

	orderData, err := o.UpdateOrderByID(orderID, &bodyBytes)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", orderID, string(bodyBytes))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		OrderDataResult{Body: orderData},
	)
}

func (o *OrderV1) OrderDeleteHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	userID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	hard := ec.QueryParam("hard")
	if hard == "true" {
		err = o.HardDeleteOrderByID(userID)
	} else {
		err = o.SoftDeleteOrderByID(userID)
	}

	if err != nil {
		hndlLog.Err(err).Msgf("DATA NOT DELETED, id %d", userID)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotDeleted(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		httpsrv.OkResult(),
	)
}
