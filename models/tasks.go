package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type TechTask struct {
	ID          int            `json:"id" db:"id"`
	Description string         `json:"object_address" db:"object_address"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (o *TechTask) SQLParamsRequest() []string {
	return []string{
		"description",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
