package models

import "github.com/sqsinformatique/rosseti-back/types"

type ElementEquipment struct {
	ID          int            `json:"id" db:"id"`
	Description string         `json:"description" db:"description"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *ElementEquipment) SQLParamsRequest() []string {
	return []string{
		"id",
		"description",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
