package dto

import (
	"time"
)

type Anl_remark struct {
	Analysis_id            string     `json:"analysis_id"`
	Anl_source             string     `json:"anl_source"`
	Remark_seq_no          int        `json:"remark_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Referenced_column_name *string    `json:"referenced_column_name"`
	Referenced_ppdm_guid   *string    `json:"referenced_ppdm_guid"`
	Referenced_system_id   *string    `json:"referenced_system_id"`
	Referenced_table_name  *string    `json:"referenced_table_name"`
	Remark                 *string    `json:"remark"`
	Remark_ba_id           *string    `json:"remark_ba_id"`
	Remark_type            *string    `json:"remark_type"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
