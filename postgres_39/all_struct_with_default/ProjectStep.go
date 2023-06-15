package dto

type Project_step struct {
	Project_id                  string   `json:"project_id" default:"source"`
	Step_id                     string   `json:"step_id" default:"source"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Actual_end_date             *string  `json:"actual_end_date" default:""`
	Actual_start_date           *string  `json:"actual_start_date" default:""`
	Actual_time_elapsed         *float64 `json:"actual_time_elapsed" default:""`
	Actual_time_elapsed_ouom    *string  `json:"actual_time_elapsed_ouom" default:""`
	Actual_time_elapsed_uom     *string  `json:"actual_time_elapsed_uom" default:""`
	Critical_date               *string  `json:"critical_date" default:""`
	Description                 *string  `json:"description" default:""`
	Due_date                    *string  `json:"due_date" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Estimated_time_elapsed      *float64 `json:"estimated_time_elapsed" default:""`
	Estimated_time_elapsed_ouom *string  `json:"estimated_time_elapsed_ouom" default:""`
	Estimated_time_elapsed_uom  *string  `json:"estimated_time_elapsed_uom" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Plan_end_date               *string  `json:"plan_end_date" default:""`
	Plan_start_date             *string  `json:"plan_start_date" default:""`
	Plan_step_id                *string  `json:"plan_step_id" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Project_plan_id             *string  `json:"project_plan_id" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Step_name                   *string  `json:"step_name" default:""`
	Step_seq_no                 *int     `json:"step_seq_no" default:""`
	Step_type                   *string  `json:"step_type" default:""`
	Where_completed             *string  `json:"where_completed" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
