package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Act struct {
	ID             int            `json:"id" db:"id"`
	StaffID        int            `json:"staff_id" db:"staff_id"`
	StaffDesc      *Profile       `json:"staff_desc,omitempty"`
	SuperviserID   int            `json:"superviser_id" db:"superviser_id"`
	SuperviserDesc *Profile       `json:"superviser_desc,omitempty"`
	ObjectID       int            `json:"object_id" db:"object_id"`
	ObjectDecs     *Object        `json:"object_desc,omitempty"`
	ReviewID       int            `json:"review_id" db:"review_id"`
	ReviewDesc     *Review        `json:"review_desc,omitempty"`
	EndAt          types.NullTime `json:"end_at" db:"end_at"`
	Finished       bool           `json:"finished" db:"finished"`
	Approved       bool           `json:"approved" db:"approved"`
	Reverted       bool           `json:"reverted" db:"reverted"`
	StaffSign      string         `json:"staff_sign" db:"staff_sign"`
	SuperviserSign string         `json:"superviser_sign" db:"superviser_sign"`
	ActDetailDesc  interface{}    `json:"act_detail_desc"`
	Meta           types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Act) SQLParamsRequest() []string {
	return []string{
		"staff_id",
		"superviser_id",
		"object_id",
		"review_id",
		"finished",
		"approved",
		"reverted",
		"staff_sign",
		"superviser_sign",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
