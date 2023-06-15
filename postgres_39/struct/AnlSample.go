package dto

import (
	"time"
)

type Anl_sample struct {
	Analysis_id         string     `json:"analysis_id"`
	Anl_source          string     `json:"anl_source"`
	Sample_id           string     `json:"sample_id"`
	Active_ind          *string    `json:"active_ind"`
	Batch_id            *string    `json:"batch_id"`
	Created_by_step_ind *string    `json:"created_by_step_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	End_date            *time.Time `json:"end_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Sample_description  *string    `json:"sample_description"`
	Source              *string    `json:"source"`
	Standard_sample_ind *string    `json:"standard_sample_ind"`
	Step_seq_no         *int       `json:"step_seq_no"`
	Used_by_step_ind    *string    `json:"used_by_step_ind"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
