package dto

import (
	"time"
)

type Project_step struct {
	Project_id                  string     `json:"project_id"`
	Step_id                     string     `json:"step_id"`
	Active_ind                  *string    `json:"active_ind"`
	Actual_end_date             *time.Time `json:"actual_end_date"`
	Actual_start_date           *time.Time `json:"actual_start_date"`
	Actual_time_elapsed         *float64   `json:"actual_time_elapsed"`
	Actual_time_elapsed_ouom    *string    `json:"actual_time_elapsed_ouom"`
	Actual_time_elapsed_uom     *string    `json:"actual_time_elapsed_uom"`
	Critical_date               *time.Time `json:"critical_date"`
	Description                 *string    `json:"description"`
	Due_date                    *time.Time `json:"due_date"`
	Effective_date              *time.Time `json:"effective_date"`
	Estimated_time_elapsed      *float64   `json:"estimated_time_elapsed"`
	Estimated_time_elapsed_ouom *string    `json:"estimated_time_elapsed_ouom"`
	Estimated_time_elapsed_uom  *string    `json:"estimated_time_elapsed_uom"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Plan_end_date               *time.Time `json:"plan_end_date"`
	Plan_start_date             *time.Time `json:"plan_start_date"`
	Plan_step_id                *string    `json:"plan_step_id"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Project_plan_id             *string    `json:"project_plan_id"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Step_name                   *string    `json:"step_name"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	Step_type                   *string    `json:"step_type"`
	Where_completed             *string    `json:"where_completed"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
