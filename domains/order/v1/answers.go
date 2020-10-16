package orderv1

import (
	// local
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/models"
)

type OrderDataResult httpsrv.ResultAnsw

type ArrayOfOrderData []models.Order
