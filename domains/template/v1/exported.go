package templatev1

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/cfg"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type empty struct{}

type TemplateV1 struct {
	log      zerolog.Logger
	cfg      *cfg.AppCfg
	mongodb  **mongo.Client
	publicV1 *echo.Group
	userV1   *userv1.UserV1
}

func NewTemplateV1(ctx *context.Context, userV1 *userv1.UserV1) (*TemplateV1, error) {
	if ctx == nil {
		return nil, errors.New("empty context or config or profilev1 client or orm client")
	}

	t := &TemplateV1{}
	t.log = ctx.GetPackageLogger(empty{})
	t.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	t.cfg = ctx.Config
	t.mongodb = ctx.GetMongoDB()
	t.userV1 = userV1

	t.publicV1.GET("/templates/:id", t.userV1.Introspect(t.templateGetHandler, types.User))
	t.publicV1.POST("/templates", t.userV1.Introspect(t.templatePostHandler, types.Admin))
	t.publicV1.PUT("/templates/:id", t.userV1.Introspect(t.templatePutHandler, types.Admin))
	t.publicV1.DELETE("/templates/:id", t.userV1.Introspect(t.templateDeleteHandler, types.Admin))

	return t, nil
}
