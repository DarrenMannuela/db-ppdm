package dto

import (
	"time"
)

type Class_system_xref struct {
	Classification_system_id  string     `json:"classification_system_id"`
	Classification_system_id2 string     `json:"classification_system_id_2"`
	Equiv_id                  string     `json:"equiv_id"`
	Active_ind                *string    `json:"active_ind"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Source_document           *string    `json:"source_document"`
	System_xref_type          *string    `json:"system_xref_type"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
