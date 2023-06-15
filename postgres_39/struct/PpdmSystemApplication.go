package dto

import (
	"time"
)

type Ppdm_system_application struct {
	System_id                 string     `json:"system_id"`
	Sw_application_id         string     `json:"sw_application_id"`
	Application_seq_no        int        `json:"application_seq_no"`
	Active_ind                *string    `json:"active_ind"`
	Application_function_desc *string    `json:"application_function_desc"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Language_id               *string    `json:"language_id"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
