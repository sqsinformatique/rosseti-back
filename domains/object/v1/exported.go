package objectv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"github.com/sqsinformatique/rosseti-back/types"
)

type empty struct{}

type ObjectV1 struct {
	log      zerolog.Logger
	db       **sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
	userV1   *userv1.UserV1
}

func NewObjectV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1) (*ObjectV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	o := &ObjectV1{}
	o.log = ctx.GetPackageLogger(empty{})
	o.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	o.db = ctx.GetDatabase()
	o.userV1 = userV1
	o.orm = orm

	o.publicV1.POST("/objects", o.userV1.Introspect(o.ObjectPostHandler, types.Admin))
	o.publicV1.GET("/objects/:id", o.userV1.Introspect(o.ObjectGetHandler, types.Admin))
	o.publicV1.GET("/objectssearch", o.userV1.Introspect(o.ObjecSearchGetHandler, types.Admin))
	o.publicV1.PUT("/objects/:id", o.userV1.Introspect(o.ObjectPutHandler, types.Admin))
	o.publicV1.DELETE("/objects/:id", o.userV1.Introspect(o.ObjectDeleteHandler, types.Admin))

	return o, nil
}
