package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Review struct {
	ID          int            `json:"id" db:"id"`
	Description string         `json:"description" db:"description"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Review) SQLParamsRequest() []string {
	return []string{
		"description",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
