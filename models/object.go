package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Object struct {
	ID      int            `json:"id" db:"id"`
	Address string         `json:"object_address" db:"object_address"`
	Name    string         `json:"object_name" db:"object_name"`
	Meta    types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (o *Object) SQLParamsRequest() []string {
	return []string{
		"object_address",
		"object_name",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
