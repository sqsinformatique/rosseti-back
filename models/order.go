package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Order struct {
	ID       int            `json:"id" db:"id"`
	ObjectID int            `json:"object_id" db:"object_id"`
	StaffID  int            `json:"staff_id" db:"staff_id"`
	Meta     types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Order) SQLParamsRequest() []string {
	return []string{
		"object_id",
		"staff_id",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
