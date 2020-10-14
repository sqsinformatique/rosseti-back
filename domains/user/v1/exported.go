package userv1

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

type UserV1 struct {
	log      zerolog.Logger
	db       *sqlx.DB
	orm      *orm.ORM
	publicV1 *echo.Group
}

func NewUserV1(ctx *context.Context, orm *orm.ORM) (*UserV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	u := &UserV1{}
	u.log = ctx.GetPackageLogger(empty{})
	u.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	u.db = ctx.GetDatabase()

	u.publicV1.POST("/user", u.userPostHandler)
	u.publicV1.GET("/users/:id", u.userGetHandler)
	u.publicV1.PUT("/users/:id", u.UserPutHandler)
	u.publicV1.PUT("/credentials/:id", u.CredsPutHandler)
	u.publicV1.POST("/credentials", u.CredsPostHandler)
	u.publicV1.DELETE("/users/:id", u.UserDeleteHandler)

	return u, nil
}
