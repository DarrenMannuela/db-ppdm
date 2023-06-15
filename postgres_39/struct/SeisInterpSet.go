package dto

import (
	"time"
)

type Seis_interp_set struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Interp_set_id      string     `json:"interp_set_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Interpreter        *string    `json:"interpreter"`
	Interp_date        *time.Time `json:"interp_date"`
	Interp_objective   *string    `json:"interp_objective"`
	Interp_set_name    *string    `json:"interp_set_name"`
	Interp_type        *string    `json:"interp_type"`
	Pick_method        *string    `json:"pick_method"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Trace_position     *string    `json:"trace_position"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
