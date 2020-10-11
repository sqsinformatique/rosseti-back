package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type User struct {
	ID    int            `json:"id" db:"id"`
	Hash  string         `json:"user_hash" db:"user_hash"`
	Login string         `json:"user_login" db:"user_login"`
	Email string         `json:"user_email" db:"user_email"`
	Phone string         `json:"user_phone" db:"user_phone"`
	Meta  types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *User) SQLParamsRequest() []string {
	return []string{
		"user_hash",
		"hash_login",
		"hash_email",
		"user_phone",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
