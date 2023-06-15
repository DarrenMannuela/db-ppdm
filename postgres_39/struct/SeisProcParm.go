package dto

import (
	"time"
)

type Seis_proc_parm struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Seis_proc_set_id   string     `json:"seis_proc_set_id"`
	Process_step_id    string     `json:"process_step_id"`
	Parameter_id       string     `json:"parameter_id"`
	Active_ind         *string    `json:"active_ind"`
	Date_format_desc   *time.Time `json:"date_format_desc"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Parameter_desc     *string    `json:"parameter_desc"`
	Parameter_origin   *string    `json:"parameter_origin"`
	Parameter_type     *string    `json:"parameter_type"`
	Parameter_uom      *string    `json:"parameter_uom"`
	Parameter_value_1  *float64   `json:"parameter_value_1"`
	Parameter_value_2  *float64   `json:"parameter_value_2"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
