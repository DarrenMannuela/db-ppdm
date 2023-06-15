package dto

type Well_drill_mud_intrvl struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Source                string   `json:"source" default:"source"`
	Media_type            string   `json:"media_type" default:"source"`
	Depth_obs_no          int      `json:"depth_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Casing_depth          *float64 `json:"casing_depth" default:""`
	Casing_depth_ouom     *string  `json:"casing_depth_ouom" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	End_date              *string  `json:"end_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Max_mud_weight        *float64 `json:"max_mud_weight" default:""`
	Max_mud_weight_ouom   *string  `json:"max_mud_weight_ouom" default:""`
	Max_mud_weight_uom    *string  `json:"max_mud_weight_uom" default:""`
	Max_weight_depth      *float64 `json:"max_weight_depth" default:""`
	Max_weight_depth_ouom *string  `json:"max_weight_depth_ouom" default:""`
	Mud_end_depth         *float64 `json:"mud_end_depth" default:""`
	Mud_end_depth_ouom    *string  `json:"mud_end_depth_ouom" default:""`
	Mud_start_depth       *float64 `json:"mud_start_depth" default:""`
	Mud_start_depth_ouom  *string  `json:"mud_start_depth_ouom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Start_date            *string  `json:"start_date" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
