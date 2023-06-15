package dto

import (
	"time"
)

type Anl_table_result struct {
	Analysis_id        string     `json:"analysis_id"`
	Anl_source         string     `json:"anl_source"`
	Ppdm_system_id     string     `json:"ppdm_system_id"`
	Ppdm_table_name    string     `json:"ppdm_table_name"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Step_seq_no        *int       `json:"step_seq_no"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
