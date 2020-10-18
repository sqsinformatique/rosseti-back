package elementequipmentv1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ElementEquipmentV1) ElementEquipmentPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var elementequipment models.ElementEquipment
	err = ec.Bind(&elementequipment)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE USER FAILED %+v", &elementequipment)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	elementequipmentData, err := o.CreateElementEquipment(&elementequipment)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE ORDER FAILED %+v", &elementequipment)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.CreateFailed(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ElementEquipmentDataResult{Body: elementequipmentData},
	)
}

func (o *ElementEquipmentV1) ElementEquipmentGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	elementequipmentID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	elementequipmentData, err := o.GetElementEquipmentByID(elementequipmentID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", elementequipmentID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ElementEquipmentDataResult{Body: elementequipmentData},
	)
}

func (o ElementEquipmentV1) ElementEquipmentSearchGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var value models.Search

	err = ec.Bind(&value)
	if err != nil {
		hndlLog.Err(err).Msg("BAD REQUEST")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	techTaskData, err := o.SearchElementEquipmentsByName(&value)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get techTaskData %+v", &value)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ElementEquipmentDataResult{Body: techTaskData},
	)
}

func (o *ElementEquipmentV1) ElementEquipmentPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	elementequipmentID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
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
			hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %d", elementequipmentID)

			return ec.JSON(
				http.StatusBadRequest,
				httpsrv.BadRequest(err),
			)
		}
	}

	elementequipmentData, err := o.UpdateElementEquipmentByID(elementequipmentID, &bodyBytes)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", elementequipmentID, string(bodyBytes))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ElementEquipmentDataResult{Body: elementequipmentData},
	)
}

func (o *ElementEquipmentV1) ElementEquipmentDeleteHandler(ec echo.Context) (err error) {
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
		err = o.HardDeleteElementEquipmentByID(userID)
	} else {
		err = o.SoftDeleteElementEquipmentByID(userID)
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
