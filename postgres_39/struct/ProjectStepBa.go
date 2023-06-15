package dto

import (
	"time"
)

type Project_step_ba struct {
	Project_id            string     `json:"project_id"`
	Business_associate_id string     `json:"business_associate_id"`
	Role                  string     `json:"role"`
	Role_seq_no           int        `json:"role_seq_no"`
	Step_id               string     `json:"step_id"`
	Step_ba_obs_no        int        `json:"step_ba_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Actual_ind            *string    `json:"actual_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Plan_ind              *string    `json:"plan_ind"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
