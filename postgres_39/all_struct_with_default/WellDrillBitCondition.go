package dto

type Well_drill_bit_condition struct {
	Uwi                 string  `json:"uwi" default:"source"`
	Bit_source          string  `json:"bit_source" default:"source"`
	Bit_interval_obs_no int     `json:"bit_interval_obs_no" default:"1"`
	Condition_obs_no    int     `json:"condition_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Bearing_condition   *string `json:"bearing_condition" default:""`
	Drill_bit_condition *string `json:"drill_bit_condition" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
