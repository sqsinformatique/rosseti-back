package actv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	"github.com/sqsinformatique/rosseti-back/internal/cfg"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"go.mongodb.org/mongo-driver/mongo"
)

type empty struct{}

type ActV1 struct {
	log       zerolog.Logger
	cfg       *cfg.AppCfg
	db        *sqlx.DB
	mongodb   *mongo.Client
	orm       *orm.ORM
	profilev1 *profilev1.ProfileV1
	publicV1  *echo.Group
}

func NewActV1(ctx *context.Context, config *cfg.AppCfg, profilev1 *profilev1.ProfileV1, orm *orm.ORM) (*ActV1, error) {
	if ctx == nil || config == nil || profilev1 == nil || orm == nil {
		return nil, errors.New("empty context or config or profilev1 client or orm client")
	}

	p := &ActV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.profilev1 = profilev1
	p.cfg = config
	p.db = ctx.GetDatabase()
	p.mongodb = ctx.GetMongoDB()

	p.publicV1.GET("/act", p.actGetHandler)
	p.publicV1.POST("/act", p.actPostHandler)
	p.publicV1.POST("/act/:actid/images", p.actPostImagesHandler)

	return p, nil
}
