package dto

type Well_log struct {
	Uwi                string   `json:"uwi" default:"source"`
	Well_log_id        string   `json:"well_log_id" default:"source"`
	Source             string   `json:"source" default:"source"`
	Acquired_for_ba_id *string  `json:"acquired_for_ba_id" default:""`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Base_depth_ouom    *string  `json:"base_depth_ouom" default:""`
	Bypass_ind         *string  `json:"bypass_ind" default:""`
	Cased_hole_ind     *string  `json:"cased_hole_ind" default:""`
	Composite_ind      *string  `json:"composite_ind" default:""`
	Depth_type         *string  `json:"depth_type" default:""`
	Dictionary_id      *string  `json:"dictionary_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Log_job_id         *string  `json:"log_job_id" default:""`
	Log_job_source     *string  `json:"log_job_source" default:""`
	Log_ref_num        *string  `json:"log_ref_num" default:""`
	Log_title          *string  `json:"log_title" default:""`
	Log_tool_pass_no   *int     `json:"log_tool_pass_no" default:""`
	Mwd_ind            *string  `json:"mwd_ind" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Top_depth          *float64 `json:"top_depth" default:""`
	Top_depth_ouom     *string  `json:"top_depth_ouom" default:""`
	Trip_obs_no        *int     `json:"trip_obs_no" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
