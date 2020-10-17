package actsdetailv1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ActsDetailV1) ActsDetailPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var actsdetail models.ActsDetail
	err = ec.Bind(&actsdetail)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE USER FAILED %+v", &actsdetail)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actsdetailData, err := o.CreateActsDetail(&actsdetail)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE ORDER FAILED %+v", &actsdetail)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.CreateFailed(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActsDetailDataResult{Body: actsdetailData},
	)
}

func (o *ActsDetailV1) ActsDetailGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	actsdetailID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actsdetailData, err := o.GetActsDetailByID(actsdetailID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", actsdetailID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActsDetailDataResult{Body: actsdetailData},
	)
}

func (o *ActsDetailV1) ActsDetailPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	actsdetailID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
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
			hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %d", actsdetailID)

			return ec.JSON(
				http.StatusBadRequest,
				httpsrv.BadRequest(err),
			)
		}
	}

	actsdetailData, err := o.UpdateActsDetailByID(actsdetailID, &bodyBytes)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", actsdetailID, string(bodyBytes))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActsDetailDataResult{Body: actsdetailData},
	)
}

func (o *ActsDetailV1) ActsDetailDeleteHandler(ec echo.Context) (err error) {
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
		err = o.HardDeleteActsDetailByID(userID)
	} else {
		err = o.SoftDeleteActsDetailByID(userID)
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
