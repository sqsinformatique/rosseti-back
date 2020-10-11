package actv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
)

type empty struct{}

type ActV1 struct {
	log       zerolog.Logger
	db        *sqlx.DB
	orm       *orm.ORM
	profilev1 *profilev1.ProfileV1
	publicV1  *echo.Group
}

func NewActV1(ctx *context.Context, profilev1 *profilev1.ProfileV1, orm *orm.ORM) (*ActV1, error) {
	if ctx == nil || profilev1 == nil || orm == nil {
		return nil, errors.New("empty context or profilev1 client or orm client")
	}

	p := &ActV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.profilev1 = profilev1
	p.db = ctx.GetDatabase()

	p.publicV1.GET("/act", p.actGetHandler)

	return p, nil
}
