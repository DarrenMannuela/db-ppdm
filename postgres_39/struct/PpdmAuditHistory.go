package dto

import (
	"time"
)

type Ppdm_audit_history struct {
	System_id                   string     `json:"system_id"`
	Table_name                  string     `json:"table_name"`
	Column_name                 string     `json:"column_name"`
	Audit_row_guid              string     `json:"audit_row_guid"`
	Audit_seq_no                int        `json:"audit_seq_no"`
	Active_ind                  *string    `json:"active_ind"`
	Audit_authorized_by_ba_id   *string    `json:"audit_authorized_by_ba_id"`
	Audit_created_by_ba_id      *string    `json:"audit_created_by_ba_id"`
	Audit_datetime              *time.Time `json:"audit_datetime"`
	Audit_desc                  *string    `json:"audit_desc"`
	Audit_reason                *string    `json:"audit_reason"`
	Audit_source                *string    `json:"audit_source"`
	Audit_type                  *string    `json:"audit_type"`
	Audit_verified_by_ba_id     *string    `json:"audit_verified_by_ba_id"`
	Data_type                   *string    `json:"data_type"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	New_date_value              *time.Time `json:"new_date_value"`
	New_numeric_value           *float64   `json:"new_numeric_value"`
	New_numeric_value_ouom      *string    `json:"new_numeric_value_ouom"`
	New_numeric_value_uom       *string    `json:"new_numeric_value_uom"`
	New_text_value              *string    `json:"new_text_value"`
	Null_description            *string    `json:"null_description"`
	Original_date_value         *time.Time `json:"original_date_value"`
	Original_numeric_value      *float64   `json:"original_numeric_value"`
	Original_numeric_value_ouom *string    `json:"original_numeric_value_ouom"`
	Original_numeric_value_uom  *string    `json:"original_numeric_value_uom"`
	Original_text_value         *string    `json:"original_text_value"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Retention_period            *string    `json:"retention_period"`
	Source                      *string    `json:"source"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
