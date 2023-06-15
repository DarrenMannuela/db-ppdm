package dto

type Project_condition struct {
	Project_id            string  `json:"project_id" default:"source"`
	Condition_obs_no      int     `json:"condition_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Condition_type        *string `json:"condition_type" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	End_date              *string `json:"end_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Project_type          *string `json:"project_type" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Start_date            *string `json:"start_date" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
