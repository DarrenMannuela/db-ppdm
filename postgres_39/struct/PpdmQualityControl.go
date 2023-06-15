package dto

import (
	"time"
)

type Ppdm_quality_control struct {
	System_id                  string     `json:"system_id"`
	Table_name                 string     `json:"table_name"`
	Qc_seq_no                  int        `json:"qc_seq_no"`
	Active_ind                 *string    `json:"active_ind"`
	Checked_by_ba_id           *string    `json:"checked_by_ba_id"`
	Column_name                *string    `json:"column_name"`
	Current_date_value         *time.Time `json:"current_date_value"`
	Current_numeric_value      *float64   `json:"current_numeric_value"`
	Current_numeric_value_ouom *string    `json:"current_numeric_value_ouom"`
	Current_numeric_value_uom  *string    `json:"current_numeric_value_uom"`
	Current_row_guid           *string    `json:"current_row_guid"`
	Current_text_value         *string    `json:"current_text_value"`
	Data_type                  *string    `json:"data_type"`
	Done_by_ba_id              *string    `json:"done_by_ba_id"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Null_description           *string    `json:"null_description"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Qc_datetime                *time.Time `json:"qc_datetime"`
	Qc_desc                    *string    `json:"qc_desc"`
	Qc_status                  *string    `json:"qc_status"`
	Qc_type                    *string    `json:"qc_type"`
	Quality_type               *string    `json:"quality_type"`
	Remark                     *string    `json:"remark"`
	Retention_period           *string    `json:"retention_period"`
	Source                     *string    `json:"source"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
