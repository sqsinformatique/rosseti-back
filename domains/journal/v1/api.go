package journalv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/logger"
)

func (j *JournalV1) JournalGetHandler(ec echo.Context) (err error) {
	// Main code of handler
	hndlLog := logger.HandlerLogger(&j.log, ec)

	journalData, err := j.GetJournalData()
	if err != nil {
		hndlLog.Err(err).Msg("GET JOURNAL FAILED")

		return ec.JSON(
			http.StatusBadRequest,
			httpsrv.BadRequest(err),
		)
	}

	return ec.JSON(
		http.StatusOK,
		JournalDataResult{Body: journalData},
	)
}
