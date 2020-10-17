package actv1

import (
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (a *ActV1) actGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	actID := ec.Param("actid")

	actData, err := a.GetActByID(actID)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get act %s", actID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}

func (a *ActV1) actsByStaffIDGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	userID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actData, err := a.GetActsByStaffID(userID)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get act %d", userID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}

func (a *ActV1) actsBySuperviserIDGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	userID, err := strconv.ParseInt(ec.Param("id"), 10, 64)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s", ec.Param("id"))

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actData, err := a.GetActsBySuperviserID(userID)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get act %d", userID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}

func (a *ActV1) actPostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	var act models.Act
	err = ec.Bind(&act)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to unmarshal act")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actData, err := a.CreateAct(&act)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to unmarshal act")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}

func (a *ActV1) actPostImagesHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	actID := ec.Param("actid")

	multipartForm, err := ec.MultipartForm()
	if err != nil {
		hndlLog.Err(err).Msgf("failed to read multipartform")
		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	err = a.CreateImages(actID, multipartForm)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to upload images")
		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		httpsrv.OkResult(),
	)
}

func (a *ActV1) actGetImageHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	actID := ec.Param("actid")
	imageID := ec.Param("id")

	gridFile, size, err := a.GetImage(actID, imageID)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to download image")
		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	ec.Response().Header().Set("Content-Length", strconv.Itoa(int(size)))
	ec.Response().Header().Set("Content-Disposition", "inline; filename=\""+imageID+"\"")

	return ec.Stream(http.StatusOK, mime.TypeByExtension(filepath.Ext(imageID)), gridFile)
}

func (a *ActV1) ActPutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	actID := ec.Param("actid")

	var act models.Act
	err = ec.Bind(&act)
	if err != nil {
		hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %s", actID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actData, err := a.UpdateActByID(actID, &act)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s, body %+v", actID, &act)

		if err == ErrActIsFinished {
			return ec.JSON(
				http.StatusPreconditionFailed,
				httpsrv.PreconditionFailed(err),
			)
		}

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}

func (a *ActV1) ActDeleteHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	userID := ec.Param("actid")

	hard := ec.QueryParam("hard")
	if hard == "true" {
		err = a.HardDeleteActByID(userID)
	} else {
		err = a.SoftDeleteActByID(userID)
	}

	if err != nil {
		hndlLog.Err(err).Msgf("DATA NOT DELETED, id %s", userID)

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

func (a *ActV1) ActsGetByDate(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&a.log, ec)

	var timeRange models.TimeRange
	err = ec.Bind(&timeRange)
	if err != nil {
		hndlLog.Err(err).Msgf("CAN NOT GET DATA")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	actData, err := a.GetActsByDate(timeRange.TimeStart, timeRange.TimeEnd)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get act %+v", timeRange)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		ActDataResult{Body: actData},
	)
}
