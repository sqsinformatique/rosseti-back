package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Defect struct {
	ID          int            `json:"id" db:"id"`
	ElementType int            `json:"element_type" db:"element_type"`
	Description string         `json:"description" db:"description"`
	Сategory    int            `json:"category" db:"category"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Defect) SQLParamsRequest() []string {
	return []string{
		"id",
		"element_type",
		"description",
		"category",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
