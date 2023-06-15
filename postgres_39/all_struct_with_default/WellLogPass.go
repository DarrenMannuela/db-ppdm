package dto

type Well_log_pass struct {
	Uwi                string   `json:"uwi" default:"source"`
	Source             string   `json:"source" default:"source"`
	Job_id             string   `json:"job_id" default:"source"`
	Trip_obs_no        int      `json:"trip_obs_no" default:"1"`
	Log_tool_pass_no   int      `json:"log_tool_pass_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Base_depth_ouom    *string  `json:"base_depth_ouom" default:""`
	Base_log_ind       *string  `json:"base_log_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_time           *string  `json:"end_time" default:""`
	End_timezone       *string  `json:"end_timezone" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Start_time         *string  `json:"start_time" default:""`
	Start_timezone     *string  `json:"start_timezone" default:""`
	Top_depth          *float64 `json:"top_depth" default:""`
	Top_depth_ouom     *string  `json:"top_depth_ouom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
