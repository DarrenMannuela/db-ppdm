package dto

type Project_step_time struct {
	Project_id              string   `json:"project_id" default:"source"`
	Step_id                 string   `json:"step_id" default:"source"`
	Time_obs_no             int      `json:"time_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Business_associate_id   *string  `json:"business_associate_id" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	End_date                *string  `json:"end_date" default:""`
	End_time                *string  `json:"end_time" default:""`
	End_timezone            *string  `json:"end_timezone" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Plan_ind                *string  `json:"plan_ind" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Start_date              *string  `json:"start_date" default:""`
	Start_time              *string  `json:"start_time" default:""`
	Start_timezone          *string  `json:"start_timezone" default:""`
	Total_time_elapsed      *float64 `json:"total_time_elapsed" default:""`
	Total_time_elapsed_ouom *string  `json:"total_time_elapsed_ouom" default:""`
	Total_time_elapsed_uom  *string  `json:"total_time_elapsed_uom" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
