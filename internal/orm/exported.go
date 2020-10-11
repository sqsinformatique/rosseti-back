package orm

import (
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/utils"
)

var ErrDBConnNotEstablished = errors.New("database connection not established")

type Inserter interface {
	SQLParamsRequest() []string
}

type ORM struct {
	schema string
	db     *sqlx.DB
}

func NewORM(schema string, ctx *context.Context) (*ORM, error) {
	if ctx == nil || schema == "" {
		return nil, errors.New("empty context or schema")
	}

	o := &ORM{}
	o.schema = schema
	o.db = ctx.GetDatabase()

	return o, nil
}

func (o *ORM) InsertInto(target string, data Inserter) (interface{}, error) {
	if o.db == nil {
		return nil, ErrDBConnNotEstablished
	}

	stmt, err := o.db.PrepareNamed(
		o.db.Rebind(utils.JoinStrings(" ", "INSERT INTO", o.schema+"."+target, "("+strings.Join(data.SQLParamsRequest(), ", ")+")",
			"VALUES", "("+":"+strings.Join(data.SQLParamsRequest(), ", :")+") RETURNING *")))
	if err != nil {
		return nil, err
	}

	err = stmt.Get(data, data)
	stmt.Close()

	return data, err
}

func (o *ORM) Update(target string, writeData Inserter) (interface{}, error) {
	query := make([]string, 0, len(writeData.SQLParamsRequest()))
	for _, param := range writeData.SQLParamsRequest() {
		query = append(query, param+"=:"+param)
	}

	if o.db == nil {
		return nil, ErrDBConnNotEstablished
	}

	_, err := o.db.NamedExec(
		o.db.Rebind(utils.JoinStrings(" ", "UPDATE"+o.schema+"."+target+"SET", strings.Join(query, ", "),
			"WHERE id=:id")),
		writeData)
	if err != nil {
		return nil, err
	}

	return writeData, nil
}
