package actv1

import (
	"net/http"

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
		hndlLog.Err(err).Msgf("failed to connect to database images")
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
