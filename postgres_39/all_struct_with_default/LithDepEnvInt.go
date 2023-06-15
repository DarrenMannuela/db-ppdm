package dto

type Lith_dep_env_int struct {
	Lithology_log_id         string   `json:"lithology_log_id" default:"source"`
	Lith_log_source          string   `json:"lith_log_source" default:"source"`
	Depth_obs_no             int      `json:"depth_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Base_depth               *float64 `json:"base_depth" default:""`
	Base_depth_ouom          *string  `json:"base_depth_ouom" default:""`
	Depositional_environment *string  `json:"depositional_environment" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Top_depth                *float64 `json:"top_depth" default:""`
	Top_depth_ouom           *string  `json:"top_depth_ouom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
