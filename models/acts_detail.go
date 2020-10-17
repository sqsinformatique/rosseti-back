package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type ActsDetail struct {
	ActID       int            `json:"act_id" db:"act_id"`
	ElementID   int            `json:"element_id" db:"element_id"`
	ElementDesc interface{}    `json:"element_desc"`
	Defects     types.NullMeta `json:"defects" db:"defects"`
	DefectsDesc interface{}    `json:"defects_desc"`
	Category    int            `json:"category" db:"category"`
	RepairedAt  types.NullTime `json:"repaired_at" db:"repaired_at"`
	Images      types.NullMeta `json:"images" db:"images"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *ActsDetail) SQLParamsRequest() []string {
	return []string{
		"act_id",
		"element_id",
		"element_type",
		"defects",
		"category",
		"repaired_at",
		"images",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
