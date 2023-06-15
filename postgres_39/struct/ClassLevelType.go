package dto

import (
	"time"
)

type Class_level_type struct {
	Classification_system_id string     `json:"classification_system_id"`
	Level_type               string     `json:"level_type"`
	Abbreviation             *string    `json:"abbreviation"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Level_order_seq_no       *int       `json:"level_order_seq_no"`
	Long_name                *string    `json:"long_name"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Ppdm_system_id           *string    `json:"ppdm_system_id"`
	Ppdm_table_name          *string    `json:"ppdm_table_name"`
	Remark                   *string    `json:"remark"`
	Selection_criteria       *string    `json:"selection_criteria"`
	Short_name               *string    `json:"short_name"`
	Source                   *string    `json:"source"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
