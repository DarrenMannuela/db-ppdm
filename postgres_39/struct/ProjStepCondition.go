package dto

import (
	"time"
)

type Proj_step_condition struct {
	Project_id            string     `json:"project_id"`
	Project_step_id       string     `json:"project_step_id"`
	Condition_obs_no      int        `json:"condition_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Effective_date        *time.Time `json:"effective_date"`
	End_date              *time.Time `json:"end_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Project_condition_id  *string    `json:"project_condition_id"`
	Project_type          *string    `json:"project_type"`
	Remark                *string    `json:"remark"`
	Required_step_id      *string    `json:"required_step_id"`
	Source                *string    `json:"source"`
	Start_date            *time.Time `json:"start_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
