package objectsdetailv1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ObjectsDetailV1) ObjectsDetailPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	var objectsdetail models.ObjectsDetail
	err = ec.Bind(&objectsdetail)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE USER FAILED %+v", &objectsdetail)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	objectsdetailData, err := o.CreateObjectsDetail(&objectsdetail)
	if err != nil {
		hndlLog.Err(err).Msgf("CREATE ORDER FAILED %+v", &objectsdetail)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.CreateFailed(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ObjectsDetailDataResult{Body: objectsdetailData},
	)
}

func (o *ObjectsDetailV1) ObjectsDetailGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	objectID, err := strconv.ParseInt(ec.Param("objectid"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	detailID, err := strconv.ParseInt(ec.Param("detailid"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	objectsdetailData, err := o.GetObjectsDetailByID(objectID, detailID)
	if err != nil {
		hndlLog.Err(err).Msgf("NOT FOUND, id %d", objectID)

		return ec.JSON(
			http.StatusNotFound,
			httpsrv.NotFound(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ObjectsDetailDataResult{Body: objectsdetailData},
	)
}

func (o *ObjectsDetailV1) ObjectsDetailSearchGetHandler(ec echo.Context) (err error) {
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

	objectsDetailData, err := o.SearchObjectsDetailByName(&value)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get techTaskData %+v", &value)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ObjectsDetailDataResult{Body: objectsDetailData},
	)
}

func (o *ObjectsDetailV1) ObjectsDetailPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	objectID, err := strconv.ParseInt(ec.Param("objectid"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	detailID, err := strconv.ParseInt(ec.Param("detailid"), 10, 64)
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
			hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %d", objectID)

			return ec.JSON(
				http.StatusBadRequest,
				httpsrv.BadRequest(err),
			)
		}
	}

	objectsdetailData, err := o.UpdateObjectsDetailByID(objectID, detailID, &bodyBytes)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %d, body %s", objectID, string(bodyBytes))

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ObjectsDetailDataResult{Body: objectsdetailData},
	)
}

func (o *ObjectsDetailV1) ObjectsDetailDeleteHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&o.log, ec)

	objectID, err := strconv.ParseInt(ec.Param("objectid"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	detailID, err := strconv.ParseInt(ec.Param("detailid"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	hard := ec.QueryParam("hard")
	if hard == "true" {
		err = o.HardDeleteObjectsDetailByID(objectID, detailID)
	} else {
		err = o.SoftDeleteObjectsDetailByID(objectID, detailID)
	}

	if err != nil {
		hndlLog.Err(err).Msgf("DATA NOT DELETED, id %d", objectID)

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
