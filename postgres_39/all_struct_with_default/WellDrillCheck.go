package dto

type Well_drill_check struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Period_obs_no         int      `json:"period_obs_no" default:"1"`
	Drill_check_obs_no    int      `json:"drill_check_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Check_set_id          *string  `json:"check_set_id" default:""`
	Check_type            *string  `json:"check_type" default:""`
	Contractor_name       *string  `json:"contractor_name" default:""`
	Contractor_rep_ba_id  *string  `json:"contractor_rep_ba_id" default:""`
	Deficient_ind         *string  `json:"deficient_ind" default:""`
	Deficient_period      *float64 `json:"deficient_period" default:""`
	Deficient_period_ouom *string  `json:"deficient_period_ouom" default:""`
	Deficient_period_uom  *string  `json:"deficient_period_uom" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Operator_name         *string  `json:"operator_name" default:""`
	Operator_rep_ba_id    *string  `json:"operator_rep_ba_id" default:""`
	Passed_ind            *string  `json:"passed_ind" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Recorded_time         *string  `json:"recorded_time" default:""`
	Recorded_timezone     *string  `json:"recorded_timezone" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
