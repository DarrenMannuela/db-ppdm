package dto

import (
	"time"
)

type Ppdm_audit_history_rem struct {
	System_id          string     `json:"system_id"`
	Table_name         string     `json:"table_name"`
	Column_name        string     `json:"column_name"`
	Audit_row_guid     string     `json:"audit_row_guid"`
	Audit_seq_no       int        `json:"audit_seq_no"`
	Remark_seq_no      int        `json:"remark_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Audit_datetime     *time.Time `json:"audit_datetime"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Remark_about_ba_id *string    `json:"remark_about_ba_id"`
	Remark_by_ba_id    *string    `json:"remark_by_ba_id"`
	Remark_for_ba_id   *string    `json:"remark_for_ba_id"`
	Remark_type        *string    `json:"remark_type"`
	Remark_use_type    *string    `json:"remark_use_type"`
	Retention_period   *string    `json:"retention_period"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
