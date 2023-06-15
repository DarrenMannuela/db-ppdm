package dto

import (
	"time"
)

type Class_level struct {
	Classification_system_id string     `json:"classification_system_id"`
	Class_level_id           string     `json:"class_level_id"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Level_definition         *string    `json:"level_definition"`
	Level_name               *string    `json:"level_name"`
	Level_ref_num            *string    `json:"level_ref_num"`
	Level_seq_no             *int       `json:"level_seq_no"`
	Level_type               *string    `json:"level_type"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Retention_period         *string    `json:"retention_period"`
	Source                   *string    `json:"source"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
