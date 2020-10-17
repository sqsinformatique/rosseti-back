package models

import (
	"github.com/sqsinformatique/rosseti-back/types"
)

type Order struct {
	ID               int            `json:"id" db:"id"`
	ObjectID         int            `json:"object_id" db:"object_id"`
	ObjectDesc       *Object        `json:"object_desc"`
	TechTasksDesc    interface{}    `json:"tech_tasks_desk"`
	SuperviserDesc   *Profile       `json:"superviser_desc"`
	TechTasks        types.NullMeta `json:"tech_tasks" db:"tech_tasks"`
	SuperviserID     int            `json:"superviser_id" db:"superviser_id"`
	StartAt          types.NullTime `json:"start_at" db:"start_at"`
	EndAt            types.NullTime `json:"end_at" db:"end_at"`
	StaffID          int            `json:"staff_id" db:"staff_id"`
	StaffDesc        *Profile       `json:"staff_desc"`
	SuperviserSign   string         `json:"superviser_sign" db:"superviser_sign"`
	SuperviserSignEt types.NullTime `json:"superviser_sign_at" db:"superviser_sign_at"`
	StaffSign        string         `json:"staff_sign" db:"staff_sign"`
	StaffSignEt      types.NullTime `json:"staff_sign_at" db:"staff_sign_at"`
	Meta             types.NullMeta `json:"meta" db:"meta"`
	Timestamp
}

func (u *Order) SQLParamsRequest() []string {
	return []string{
		"object_id",
		"tech_tasks",
		"superviser_id",
		"start_at",
		"end_at",
		"staff_id",
		"superviser_sign",
		"superviser_sign_at",
		"staff_sign",
		"staff_sign_at",
		"meta",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
