package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Act struct {
	ID       int  `json:"id" db:"id" bson:"id"`
	UserID   int  `json:"user_id" db:"user_id" bson:"user_id"`
	Finished bool `json:"finished" db:"finished" bson:"finished"`
	// Field for document body
	Body types.NullMeta `json:"body" bson:"body"`
	Meta types.NullMeta `json:"meta" db:"meta" bson:"meta"`
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
