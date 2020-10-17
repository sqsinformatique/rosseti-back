package objectsdetailv1

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

type ObjectsDetailV1 struct {
	log      zerolog.Logger
	db       **sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
	userV1   *userv1.UserV1
}

func NewObjectsDetailV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1) (*ObjectsDetailV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	p := &ObjectsDetailV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.userV1 = userV1
	p.db = ctx.GetDatabase()
	p.orm = orm

	p.publicV1.POST("/objectsdetail", p.userV1.Introspect(p.ObjectsDetailPostHandler, types.Electrician))
	p.publicV1.GET("/objectsdetail/:id", p.userV1.Introspect(p.ObjectsDetailGetHandler, types.Electrician))
	p.publicV1.GET("/objectsdetailsearch", p.userV1.Introspect(p.ObjectsDetailSearchGetHandler, types.Electrician))
	p.publicV1.PUT("/objectsdetail/:id", p.userV1.Introspect(p.ObjectsDetailPutHandler, types.Electrician))
	p.publicV1.DELETE("/objectsdetail/:id", p.userV1.Introspect(p.ObjectsDetailDeleteHandler, types.Admin))

	return p, nil
}
