package elementequipmentv1

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

type ElementEquipmentV1 struct {
	log      zerolog.Logger
	db       **sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
	userV1   *userv1.UserV1
}

func NewElementEquipmentV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1) (*ElementEquipmentV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	p := &ElementEquipmentV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.userV1 = userV1
	p.db = ctx.GetDatabase()
	p.orm = orm

	p.publicV1.POST("/elementequipment", p.userV1.Introspect(p.ElementEquipmentPostHandler, types.Electrician))
	p.publicV1.GET("/elementequipment/:id", p.userV1.Introspect(p.ElementEquipmentGetHandler, types.Electrician))
	p.publicV1.GET("/elementequipmentsearch", p.userV1.Introspect(p.ElementEquipmentSearchGetHandler, types.Electrician))
	p.publicV1.PUT("/elementequipment/:id", p.userV1.Introspect(p.ElementEquipmentPutHandler, types.Electrician))
	p.publicV1.DELETE("/elementequipment/:id", p.userV1.Introspect(p.ElementEquipmentDeleteHandler, types.Admin))

	return p, nil
}
