package dto

import (
	"time"
)

type Seis_proc_plan_step struct {
	Seis_proc_plan_id  string     `json:"seis_proc_plan_id"`
	Proc_plan_step_id  string     `json:"proc_plan_step_id"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Step_name          *string    `json:"step_name"`
	Step_seq_no        *int       `json:"step_seq_no"`
	Step_type          *string    `json:"step_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
