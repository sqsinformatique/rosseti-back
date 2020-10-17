package techtaskv1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *TechTaskV1) TechTaskPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var techtask models.TechTask
	err = ec.Bind(&techtask)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE USER FAILED %+v", &techtask)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	techtaskData, err := o.CreateTechTask(&techtask)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE ORDER FAILED %+v", &techtask)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.CreateFailed(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TechTaskDataResult{Body: techtaskData},
	)
}

func (o *TechTaskV1) TechTaskGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	techtaskID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	techtaskData, err := o.GetTechTaskByID(techtaskID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", techtaskID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TechTaskDataResult{Body: techtaskData},
	)
}

func (o *TechTaskV1) TechSearchGetHandler(ec echo.Context) (err error) {
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

	techTaskData, err := o.SearchTechTaskByName(&value)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get techTaskData %+v", &value)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TechTaskDataResult{Body: techTaskData},
	)
}

func (o *TechTaskV1) TechTaskPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	techtaskID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
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
			hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %d", techtaskID)

			return ec.JSON(
				http.StatusBadRequest,
				httpsrv.BadRequest(err),
			)
		}
	}

	techtaskData, err := o.UpdateTechTaskByID(techtaskID, &bodyBytes)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", techtaskID, string(bodyBytes))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TechTaskDataResult{Body: techtaskData},
	)
}

func (o *TechTaskV1) TechTaskDeleteHandler(ec echo.Context) (err error) {
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
		err = o.HardDeleteTechTaskByID(userID)
	} else {
		err = o.SoftDeleteTechTaskByID(userID)
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
