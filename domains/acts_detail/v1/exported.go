package actsdetailv1

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	defectv1 "github.com/sqsinformatique/rosseti-back/domains/defect/v1"
	objectsdetailv1 "github.com/sqsinformatique/rosseti-back/domains/objects_detail/v1"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/orm"
	"github.com/sqsinformatique/rosseti-back/types"
)

type empty struct{}

type ActsDetailV1 struct {
	log             zerolog.Logger
	db              **sqlx.DB
	orm             *orm.ORM
	publicV1        *echo.Group
	userV1          *userv1.UserV1
	defectV1        *defectv1.DefectV1
	objectsdetailV1 *objectsdetailv1.ObjectsDetailV1
}

func NewActsDetailV1(ctx *context.Context, orm *orm.ORM, userV1 *userv1.UserV1, defectV1 *defectv1.DefectV1, objectsdetailV1 *objectsdetailv1.ObjectsDetailV1) (*ActsDetailV1, error) {
	if ctx == nil || orm == nil {
		return nil, errors.New("empty context or orm client")
	}

	p := &ActsDetailV1{}
	p.log = ctx.GetPackageLogger(empty{})
	p.publicV1 = ctx.GetHTTPGroup(httpsrv.PublicSrv, httpsrv.V1)
	p.userV1 = userV1
	p.db = ctx.GetDatabase()
	p.orm = orm
	p.defectV1 = defectV1
	p.objectsdetailV1 = objectsdetailV1

	p.publicV1.POST("/actsdetail", p.userV1.Introspect(p.ActsDetailPostHandler, types.Electrician))
	p.publicV1.GET("/actsdetail/:id", p.userV1.Introspect(p.ActsDetailGetHandler, types.Electrician))
	p.publicV1.PUT("/actsdetail/:id", p.userV1.Introspect(p.ActsDetailPutHandler, types.Electrician))
	p.publicV1.DELETE("/actsdetail/:id", p.userV1.Introspect(p.ActsDetailDeleteHandler, types.Admin))

	return p, nil
}
