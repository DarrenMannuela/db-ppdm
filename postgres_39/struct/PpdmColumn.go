package dto

import (
	"time"
)

type Ppdm_column struct {
	System_id                 string     `json:"system_id"`
	Table_name                string     `json:"table_name"`
	Column_name               string     `json:"column_name"`
	Active_ind                *string    `json:"active_ind"`
	Column_comment            *string    `json:"column_comment"`
	Column_key_method         *string    `json:"column_key_method"`
	Column_precision          *float64   `json:"column_precision"`
	Column_scale              *float64   `json:"column_scale"`
	Column_sequence           *int       `json:"column_sequence"`
	Column_size               *float64   `json:"column_size"`
	Control_column_ind        *string    `json:"control_column_ind"`
	Control_column_name       *string    `json:"control_column_name"`
	Data_type                 *string    `json:"data_type"`
	Default_ind_value         *string    `json:"default_ind_value"`
	Default_uom_id            *string    `json:"default_uom_id"`
	Default_value_method      *string    `json:"default_value_method"`
	Distinct_value_count      *float64   `json:"distinct_value_count"`
	Distinct_value_count_date *time.Time `json:"distinct_value_count_date"`
	Domain                    *string    `json:"domain"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Extension_ind             *string    `json:"extension_ind"`
	Last_system_key           *string    `json:"last_system_key"`
	Nullable_ind              *string    `json:"nullable_ind"`
	Ouom_column_name          *string    `json:"ouom_column_name"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Primary_key_sequence      *int       `json:"primary_key_sequence"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Surrogate_ind             *string    `json:"surrogate_ind"`
	Uom_column_name           *string    `json:"uom_column_name"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
