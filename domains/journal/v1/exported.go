package journalv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	actv1 "github.com/sqsinformatique/rosseti-back/domains/act/v1"
	categoryv1 "github.com/sqsinformatique/rosseti-back/domains/category/v1"
	elementequipmentv1 "github.com/sqsinformatique/rosseti-back/domains/element_equpment/v1"
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
	log                zerolog.Logger
	db                 **sqlx.DB
	orm                *orm.ORM
	publicV1           *echo.Group
	userV1             *userv1.UserV1
	objectV1           *objectv1.ObjectV1
	profileV1          *profilev1.ProfileV1
	actV1              *actv1.ActV1
	categoryV1         *categoryv1.CategoryV1
	elementequipmentV1 *elementequipmentv1.ElementEquipmentV1
}

func NewJournalV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1,
	objectV1 *objectv1.ObjectV1, profileV1 *profilev1.ProfileV1,
	actV1 *actv1.ActV1,
	categoryV1 *categoryv1.CategoryV1,
	elementequipmentV1 *elementequipmentv1.ElementEquipmentV1) (*JournalV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	j := &JournalV1{}
	j.log = ctx.GetPackageLogger(empty{})
	j.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	j.db = ctx.GetDatabase()
	j.userV1 = userV1
	j.objectV1 = objectV1
	j.profileV1 = profileV1
	j.actV1 = actV1
	j.categoryV1 = categoryV1
	j.elementequipmentV1 = elementequipmentV1
	j.orm = orm

	j.publicV1.GET("/journal", j.userV1.Introspect(j.JournalGetHandler, types.Electrician))

	return j, nil
}
