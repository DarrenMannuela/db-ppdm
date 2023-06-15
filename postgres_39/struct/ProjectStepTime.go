package dto

import (
	"time"
)

type Project_step_time struct {
	Project_id              string     `json:"project_id"`
	Step_id                 string     `json:"step_id"`
	Time_obs_no             int        `json:"time_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Business_associate_id   *string    `json:"business_associate_id"`
	Effective_date          *time.Time `json:"effective_date"`
	End_date                *time.Time `json:"end_date"`
	End_time                *time.Time `json:"end_time"`
	End_timezone            *string    `json:"end_timezone"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Plan_ind                *string    `json:"plan_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Start_date              *time.Time `json:"start_date"`
	Start_time              *time.Time `json:"start_time"`
	Start_timezone          *string    `json:"start_timezone"`
	Total_time_elapsed      *float64   `json:"total_time_elapsed"`
	Total_time_elapsed_ouom *string    `json:"total_time_elapsed_ouom"`
	Total_time_elapsed_uom  *string    `json:"total_time_elapsed_uom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
