package models

import "github.com/sqsinformatique/rosseti-back/types"

type JournalItem struct {
	FindAd            types.NullData `json:"find_at"`
	ObjectDesc        *Object        `json:"object_desc"`
	ObjectsDetailDesc *ObjectsDetail `json:"objects_detail_desc"`
	ElementTypeDesk   *ElementType   `json:"element_type_desc"`
	DefectsDesc       interface{}    `json:"defects_desc"`
	StaffDesc         *User          `json:"staff_desc"`
	Category          int            `json:"category" db:"category"`
	RapairAt          types.NullData `json:"repair_at" db:"repair_at"`
}
