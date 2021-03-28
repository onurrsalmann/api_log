package entity

import "time"

type Log struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	MethodType string `gorm:"type:varchar(255)" json:"method_type"`
	ResponseMS string `gorm:"type:varchar(255)" json:"response_ms"`
	Timestamp time.Time
}