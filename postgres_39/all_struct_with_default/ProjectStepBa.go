package dto

type Project_step_ba struct {
	Project_id            string  `json:"project_id" default:"source"`
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Role                  string  `json:"role" default:"source"`
	Role_seq_no           int     `json:"role_seq_no" default:"1"`
	Step_id               string  `json:"step_id" default:"source"`
	Step_ba_obs_no        int     `json:"step_ba_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Actual_ind            *string `json:"actual_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Plan_ind              *string `json:"plan_ind" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
