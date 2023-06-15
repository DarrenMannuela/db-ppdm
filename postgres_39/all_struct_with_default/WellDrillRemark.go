package dto

type Well_drill_remark struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Remark_seq_no       int      `json:"remark_seq_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Base_depth          *float64 `json:"base_depth" default:""`
	Base_depth_ouom     *string  `json:"base_depth_ouom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	End_period_obs_no   *int     `json:"end_period_obs_no" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Remark_by_ba_id     *string  `json:"remark_by_ba_id" default:""`
	Remark_date         *string  `json:"remark_date" default:""`
	Remark_type         *string  `json:"remark_type" default:""`
	Source              *string  `json:"source" default:""`
	Start_period_obs_no *int     `json:"start_period_obs_no" default:""`
	Strat_name_set_id   *string  `json:"strat_name_set_id" default:""`
	Strat_unit_id       *string  `json:"strat_unit_id" default:""`
	Top_depth           *float64 `json:"top_depth" default:""`
	Top_depth_ouom      *string  `json:"top_depth_ouom" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
