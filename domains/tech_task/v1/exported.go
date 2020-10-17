package techtaskv1

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

type TechTaskV1 struct {
	log      zerolog.Logger
	db       **sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
	userV1   *userv1.UserV1
}

func NewTechTaskV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1) (*TechTaskV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	p := &TechTaskV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.userV1 = userV1
	p.db = ctx.GetDatabase()
	p.orm = orm

	p.publicV1.POST("/techtask", p.userV1.Introspect(p.TechTaskPostHandler, types.Electrician))
	p.publicV1.GET("/techtask/:id", p.userV1.Introspect(p.TechTaskGetHandler, types.Electrician))
	p.publicV1.PUT("/techtask/:id", p.userV1.Introspect(p.TechTaskPutHandler, types.Electrician))
	p.publicV1.DELETE("/techtask/:id", p.userV1.Introspect(p.TechTaskDeleteHandler, types.Admin))

	return p, nil
}
