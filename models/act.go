package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Act struct {
	ID       int            `json:"id" db:"id"`
	UserID   int            `json:"user_id" db:"user_id"`
	Finished bool           `json:"finished" db:"finished"`
	Meta     types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Act) SQLParamsRequest() []string {
	return []string{
		"user_id",
		"finished",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
