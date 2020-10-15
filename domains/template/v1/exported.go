package templatev1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sqsinformatique/rosseti-back/internal/cfg"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"go.mongodb.org/mongo-driver/mongo"
)

type empty struct{}

type TemplateV1 struct {
	log      zerolog.Logger
	cfg      *cfg.AppCfg
	db       *sqlx.DB
	mongodb  *mongo.Client
	publicV1 *echo.Group
}

func NewActV1(ctx *context.Context, config *cfg.AppCfg, orm *orm.ORM) (*TemplateV1, error) {
	if ctx == nil || config == nil || orm == nil {
		return nil, errors.New("empty context or config or profilev1 client or orm client")
	}

	t := &TemplateV1{}
	t.log = ctx.GetPackageLogger(empty{})
	t.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	t.cfg = config
	t.db = ctx.GetDatabase()
	t.mongodb = ctx.GetMongoDB()

	t.publicV1.GET("/templates/:id", t.templateGetHandler)
	t.publicV1.POST("/templates", t.templatePostHandler)
	t.publicV1.PUT("/templates/:id", t.templatePutHandler)
	t.publicV1.DELETE("/templates/:id", t.templateDeleteHandler)

	return t, nil
}
