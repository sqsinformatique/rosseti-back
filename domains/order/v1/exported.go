package orderv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
)

type empty struct{}

type OrderV1 struct {
	log      zerolog.Logger
	db       *sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
}

func NewOrderV1(ctx *context.Context, orm *orm.ORM) (*OrderV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	o := &OrderV1{}
	o.log = ctx.GetPackageLogger(empty{})
	o.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	o.db = ctx.GetDatabase()

	o.publicV1.POST("/orders", o.OrderPostHandler)
	o.publicV1.GET("/orders/:id", o.OrderGetHandler)
	o.publicV1.PUT("/orders/:id", o.OrderPutHandler)
	o.publicV1.DELETE("/orders/:id", o.OrderDeleteHandler)

	return o, nil
}
