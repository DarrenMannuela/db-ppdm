package dto

import (
	"time"
)

type Class_level_alias struct {
	Classification_system_id string     `json:"classification_system_id"`
	Class_level_id           string     `json:"class_level_id"`
	Class_level_alias_id     string     `json:"class_level_alias_id"`
	Abbreviation             *string    `json:"abbreviation"`
	Active_ind               *string    `json:"active_ind"`
	Alias_long_name          *string    `json:"alias_long_name"`
	Alias_reason_type        *string    `json:"alias_reason_type"`
	Alias_short_name         *string    `json:"alias_short_name"`
	Alias_type               *string    `json:"alias_type"`
	Alias_version            *string    `json:"alias_version"`
	Amended_date             *time.Time `json:"amended_date"`
	Column_name              *string    `json:"column_name"`
	Created_date             *time.Time `json:"created_date"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Language_id              *string    `json:"language_id"`
	Original_ind             *string    `json:"original_ind"`
	Owner_ba_id              *string    `json:"owner_ba_id"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Preferred_ind            *string    `json:"preferred_ind"`
	Reason_desc              *string    `json:"reason_desc"`
	Referenced_guid          *string    `json:"referenced_guid"`
	Remark                   *string    `json:"remark"`
	Schema_entity_id         *string    `json:"schema_entity_id"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	Struckoff_date           *time.Time `json:"struckoff_date"`
	Sw_application_id        *string    `json:"sw_application_id"`
	System_id                *string    `json:"system_id"`
	Table_name               *string    `json:"table_name"`
	Use_rule                 *string    `json:"use_rule"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
