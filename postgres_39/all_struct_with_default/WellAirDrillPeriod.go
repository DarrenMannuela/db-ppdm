package dto

type Well_air_drill_period struct {
	Uwi                string   `json:"uwi" default:"source"`
	Interval_source    string   `json:"interval_source" default:"source"`
	Air_drill_obs_no   int      `json:"air_drill_obs_no" default:"1"`
	Depth_obs_no       int      `json:"depth_obs_no" default:"1"`
	Period_obs_no      int      `json:"period_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Base_depth_ouom    *string  `json:"base_depth_ouom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
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
