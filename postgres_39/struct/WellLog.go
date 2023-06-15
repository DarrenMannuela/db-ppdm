package dto

import (
	"time"
)

type Well_log struct {
	Uwi                string     `json:"uwi"`
	Well_log_id        string     `json:"well_log_id"`
	Source             string     `json:"source"`
	Acquired_for_ba_id *string    `json:"acquired_for_ba_id"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Base_depth_ouom    *string    `json:"base_depth_ouom"`
	Bypass_ind         *string    `json:"bypass_ind"`
	Cased_hole_ind     *string    `json:"cased_hole_ind"`
	Composite_ind      *string    `json:"composite_ind"`
	Depth_type         *string    `json:"depth_type"`
	Dictionary_id      *string    `json:"dictionary_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Log_job_id         *string    `json:"log_job_id"`
	Log_job_source     *string    `json:"log_job_source"`
	Log_ref_num        *string    `json:"log_ref_num"`
	Log_title          *string    `json:"log_title"`
	Log_tool_pass_no   *int       `json:"log_tool_pass_no"`
	Mwd_ind            *string    `json:"mwd_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Top_depth          *float64   `json:"top_depth"`
	Top_depth_ouom     *string    `json:"top_depth_ouom"`
	Trip_obs_no        *int       `json:"trip_obs_no"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
