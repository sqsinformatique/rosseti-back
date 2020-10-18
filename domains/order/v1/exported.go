package orderv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	objectv1 "github.com/sqsinformatique/rosseti-back/domains/object/v1"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	techtaskv1 "github.com/sqsinformatique/rosseti-back/domains/tech_task/v1"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"github.com/sqsinformatique/rosseti-back/types"
)

type empty struct{}

type OrderV1 struct {
	log        zerolog.Logger
	db         **sqlx.DB
	orm        *orm.ORM
	publicV1   *echo.Group
	userV1     *userv1.UserV1
	objectV1   *objectv1.ObjectV1
	profileV1  *profilev1.ProfileV1
	techtaskV1 *techtaskv1.TechTaskV1
}

func NewOrderV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1, objectV1 *objectv1.ObjectV1, profileV1 *profilev1.ProfileV1,
	techtaskV1 *techtaskv1.TechTaskV1) (*OrderV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	o := &OrderV1{}
	o.log = ctx.GetPackageLogger(empty{})
	o.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	o.db = ctx.GetDatabase()
	o.userV1 = userV1
	o.objectV1 = objectV1
	o.profileV1 = profileV1
	o.techtaskV1 = techtaskV1
	o.orm = orm

	o.publicV1.POST("/orders", o.userV1.Introspect(o.OrderPostHandler, types.Master))
	o.publicV1.GET("/orders/:id", o.userV1.Introspect(o.OrderGetHandler, types.Electrician))
	o.publicV1.POST("/orders/:id/signsupervisor", o.userV1.Introspect(o.OrderSignSuperviserPostHandler, types.Master))
	o.publicV1.POST("/orders/:id/signstaff", o.userV1.Introspect(o.OrderSignStaffPostHandler, types.Electrician))
	o.publicV1.GET("/orders/user/:id", o.userV1.Introspect(o.OrdersGetByUserIDHandler, types.Electrician))
	o.publicV1.PUT("/orders/:id", o.userV1.Introspect(o.OrderPutHandler, types.Master))
	o.publicV1.DELETE("/orders/:id", o.userV1.Introspect(o.OrderDeleteHandler, types.Master))

	return o, nil
}
