package objectv1

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

type ObjectV1 struct {
	log      zerolog.Logger
	db       *sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
}

func NewObjectV1(ctx *context.Context, orm *orm.ORM) (*ObjectV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	o := &ObjectV1{}
	o.log = ctx.GetPackageLogger(empty{})
	o.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	o.db = ctx.GetDatabase()

	o.publicV1.POST("/objects", o.ObjectPostHandler)
	o.publicV1.GET("/objects/:id", o.ObjectGetHandler)
	o.publicV1.PUT("/objects/:id", o.ObjectPutHandler)
	o.publicV1.DELETE("/objects/:id", o.ObjectDeleteHandler)

	return o, nil
}
