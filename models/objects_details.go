package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type ObjectsDetail struct {
	ObjectID    int            `json:"object_id" db:"object_id"`
	ElementID   int            `json:"element_id" db:"element_id"`
	ElementName string         `json:"element_name" db:"element_name"`
	ElementType int            `json:"element_type" db:"element_type"`
	Meta        types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *ObjectsDetail) SQLParamsRequest() []string {
	return []string{
		"object_id",
		"element_id",
		"element_name",
		"element_type",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
