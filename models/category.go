package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Category struct {
	ID           int            `json:"id" db:"id"`
	RapairPeriod int            `json:"rapair_period" db:"rapair_period"`
	Description  string         `json:"description" db:"description"`
	Meta         types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Category) SQLParamsRequest() []string {
	return []string{
		"id",
		"rapair_period",
		"description",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
