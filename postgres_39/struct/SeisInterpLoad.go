package dto

import (
	"time"
)

type Seis_interp_load struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Interp_set_id      string     `json:"interp_set_id"`
	Process_step_id    string     `json:"process_step_id"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Interp_origin_id   *string    `json:"interp_origin_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Process_status     *string    `json:"process_status"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_date         *time.Time `json:"start_date"`
	Step_name          *string    `json:"step_name"`
	Step_seq_no        *int       `json:"step_seq_no"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
