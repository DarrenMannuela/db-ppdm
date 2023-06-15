package dto

type Anl_gas_heat struct {
	Analysis_id             string   `json:"analysis_id" default:"source"`
	Anl_source              string   `json:"anl_source" default:"source"`
	Heat_content_obs_no     int      `json:"heat_content_obs_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Calculated_ind          *string  `json:"calculated_ind" default:""`
	Calculate_method_id     *string  `json:"calculate_method_id" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Heat_content_method     *string  `json:"heat_content_method" default:""`
	Heat_content_press      *float64 `json:"heat_content_press" default:""`
	Heat_content_press_ouom *string  `json:"heat_content_press_ouom" default:""`
	Heat_content_temp       *float64 `json:"heat_content_temp" default:""`
	Heat_content_temp_ouom  *string  `json:"heat_content_temp_ouom" default:""`
	Heat_content_value      *float64 `json:"heat_content_value" default:""`
	Heat_content_value_ouom *string  `json:"heat_content_value_ouom" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Problem_ind             *string  `json:"problem_ind" default:""`
	Remark                  *string  `json:"remark" default:""`
	Reported_ind            *string  `json:"reported_ind" default:""`
	Source                  *string  `json:"source" default:""`
	Step_seq_no             *int     `json:"step_seq_no" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
