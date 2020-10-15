package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Object struct {
	ID      int    `json:"id" db:"id"`
	Address string `json:"object_address" db:"object_address"`
	Type    string `json:"object_type" db:"object_type"`
	// Field for document body
	Meta types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (o *Object) SQLParamsRequest() []string {
	return []string{
		"object_address",
		"object_type",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
