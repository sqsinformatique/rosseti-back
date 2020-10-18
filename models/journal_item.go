package models

import "github.com/sqsinformatique/rosseti-back/types"

type JournalItem struct {
	FindAd          types.NullTime `json:"find_at"`
	ObjectDesc      *Object        `json:"object_desc"`
	ElementDesc     interface{}    `json:"element_desc"`
	ElementTypeDesk interface{}    `json:"element_type_desc"`
	DefectsDesc     interface{}    `json:"defects_desc"`
	StaffDesc       interface{}    `json:"staff_desc"`
	Category        int            `json:"category" db:"category"`
	RapairAt        types.NullTime `json:"repair_at" db:"repair_at"`
}
