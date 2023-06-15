package dto

import (
	"time"
)

type Class_system struct {
	Classification_system_id string     `json:"classification_system_id"`
	Active_ind               *string    `json:"active_ind"`
	Class_dimension          *string    `json:"class_dimension"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	System_definition        *string    `json:"system_definition"`
	System_name              *string    `json:"system_name"`
	System_owner             *string    `json:"system_owner"`
	System_ref_num           *string    `json:"system_ref_num"`
	System_version           *string    `json:"system_version"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
