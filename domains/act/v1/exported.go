package actv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/cfg"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"github.com/sqsinformatique/rosseti-back/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type empty struct{}

type ActV1 struct {
	log       zerolog.Logger
	cfg       *cfg.AppCfg
	db        **sqlx.DB
	mongodb   **mongo.Client
	orm       *orm.ORM
	profilev1 *profilev1.ProfileV1
	publicV1  *echo.Group
	userV1    *userv1.UserV1
}

func NewActV1(ctx *context.Context, profilev1 *profilev1.ProfileV1, orm *orm.ORM, userV1 *userv1.UserV1) (*ActV1, error) {
	if ctx == nil || profilev1 == nil || orm == nil {
		return nil, errors.New("empty context or profilev1 client or orm client")
	}

	a := &ActV1{}
	a.log = ctx.GetPackageLogger(empty{})
	a.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	a.profilev1 = profilev1
	a.cfg = ctx.Config
	a.db = ctx.GetDatabase()
	a.mongodb = ctx.GetMongoDB()
	a.userV1 = userV1
	a.orm = orm

	a.publicV1.GET("/acts/:actid", a.userV1.Introspect(a.actGetHandler, types.User))
	a.publicV1.GET("/acts/user/:id", a.userV1.Introspect(a.actsByUserIDGetHandler, types.User))
	a.publicV1.POST("/acts", a.userV1.Introspect(a.actPostHandler, types.User))
	a.publicV1.PUT("/acts/:actid", a.userV1.Introspect(a.ActPutHandler, types.User))
	a.publicV1.POST("/acts/:actid/images", a.userV1.Introspect(a.actPostImagesHandler, types.User))
	a.publicV1.GET("/acts/:actid/images/:id", a.userV1.Introspect(a.actGetImageHandler, types.User))
	a.publicV1.DELETE("/acts/:actid", a.userV1.Introspect(a.ActDeleteHandler, types.Master))

	return a, nil
}
