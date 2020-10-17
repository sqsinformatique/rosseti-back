package journalv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	actv1 "github.com/sqsinformatique/rosseti-back/domains/act/v1"
	objectv1 "github.com/sqsinformatique/rosseti-back/domains/object/v1"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"github.com/sqsinformatique/rosseti-back/types"
)

type empty struct{}

type JournalV1 struct {
	log       zerolog.Logger
	db        **sqlx.DB
	orm       *orm.ORM
	publicV1  *echo.Group
	userV1    *userv1.UserV1
	objectV1  *objectv1.ObjectV1
	profileV1 *profilev1.ProfileV1
	actV1     *actv1.ActV1
}

func NewJournalV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1,
	objectV1 *objectv1.ObjectV1, profileV1 *profilev1.ProfileV1,
	actV1 *actv1.ActV1) (*JournalV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	o := &JournalV1{}
	o.log = ctx.GetPackageLogger(empty{})
	o.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	o.db = ctx.GetDatabase()
	o.userV1 = userV1
	o.objectV1 = objectV1
	o.profileV1 = profileV1
	o.actV1 = actV1
	o.orm = orm

	o.publicV1.GET("/journal", o.userV1.Introspect(o.JournalGetHandler, types.Electrician))

	return o, nil
}
