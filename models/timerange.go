package models

import "github.com/sqsinformatique/rosseti-back/types"

type TimeRange struct {
	TimeStart types.NullTime `json:"time_start"`
	TimeEnd   types.NullTime `json:"time_end"`
}
