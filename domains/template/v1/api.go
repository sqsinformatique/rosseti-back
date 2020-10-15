package templatev1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (t *TemplateV1) templateGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&t.log, ec)

	actID := ec.Param("actid")

	templateData, err := t.GetTemplateByID(actID)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to get template %s", actID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TemplateDataResult{Body: templateData},
	)
}

func (t *TemplateV1) templatePostHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&t.log, ec)

	var template models.Template
	err = ec.Bind(&template)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to unmarshal act")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	templateData, err := t.CreateTemplate(&template)
	if err != nil {
		hndlLog.Err(err).Msgf("failed to unmarshal act")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TemplateDataResult{Body: templateData},
	)
}

func (t *TemplateV1) templatePutHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&t.log, ec)

	actID := ec.Param("id")

	var template models.Template
	err = ec.Bind(&template)
	if err != nil {
		hndlLog.Err(err).Msgf("ORDER DATA NOT UPDATED, id %s", actID)

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	templateData, err := t.UpdateTemplateByID(actID, &template)
	if err != nil {
		hndlLog.Err(err).Msgf("BAD REQUEST, id %s, body %+v", actID, &template)

		return ec.JSON(
			http.StatusConflict,
			httpsrv.NotUpdated(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		TemplateDataResult{Body: templateData},
	)
}

func (t *TemplateV1) templateDeleteHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&t.log, ec)

	templateID := ec.Param("id")

	hard := ec.QueryParam("hard")
	if hard == "true" {
		err = t.HardDeleteTemplateByID(templateID)
	} else {
		err = t.SoftDeleteTemplateByID(templateID)
	}

	if err != nil {
		hndlLog.Err(err).Msgf("DATA NOT DELETED, id %s", templateID)

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
