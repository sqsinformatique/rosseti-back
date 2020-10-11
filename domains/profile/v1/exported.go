package profilev1

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

type ProfileV1 struct {
	log      zerolog.Logger
	db       *sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
}

func NewProfileV1(ctx *context.Context, orm *orm.ORM) (*ProfileV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	p := &ProfileV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.db = ctx.GetDatabase()

	p.publicV1.GET("/profile", p.profileGetHandler)

	return p, nil
}
