package dto

import (
	"time"
)

type Ppdm_constraint struct {
	System_id                  string     `json:"system_id"`
	Table_name                 string     `json:"table_name"`
	Constraint_name            string     `json:"constraint_name"`
	Active_ind                 *string    `json:"active_ind"`
	Constraint_full_name       *string    `json:"constraint_full_name"`
	Constraint_type            *string    `json:"constraint_type"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Extension_ind              *string    `json:"extension_ind"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Referenced_constraint_name *string    `json:"referenced_constraint_name"`
	Referenced_system_id       *string    `json:"referenced_system_id"`
	Referenced_table_name      *string    `json:"referenced_table_name"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
