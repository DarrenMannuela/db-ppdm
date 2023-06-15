package dto

import (
	"time"
)

type Ppdm_table_history struct {
	System_id           string     `json:"system_id"`
	Table_name          string     `json:"table_name"`
	History_seq_no      int        `json:"history_seq_no"`
	Active_ind          *string    `json:"active_ind"`
	Audit_reason        *string    `json:"audit_reason"`
	Authorized_by_ba_id *string    `json:"authorized_by_ba_id"`
	Delete_record       *string    `json:"delete_record"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Retention_period    *string    `json:"retention_period"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
