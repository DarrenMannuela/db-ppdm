package dto

type Well_drill_mud_weight struct {
	Uwi                string   `json:"uwi" default:"source"`
	Source             string   `json:"source" default:"source"`
	Depth_obs_no       int      `json:"depth_obs_no" default:"1"`
	Media_obs_no       int      `json:"media_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_date           *string  `json:"end_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Mud_depth          *float64 `json:"mud_depth" default:""`
	Mud_depth_ouom     *string  `json:"mud_depth_ouom" default:""`
	Mud_weight         *float64 `json:"mud_weight" default:""`
	Mud_weight_ouom    *string  `json:"mud_weight_ouom" default:""`
	Mud_weight_uom     *string  `json:"mud_weight_uom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Reported_tvd       *float64 `json:"reported_tvd" default:""`
	Reported_tvd_ouom  *string  `json:"reported_tvd_ouom" default:""`
	Start_date         *string  `json:"start_date" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
