package dto

type Anl_gas_analysis struct {
	Analysis_id                string   `json:"analysis_id" default:"source"`
	Anl_source                 string   `json:"anl_source" default:"source"`
	Gas_anl_obs_no             int      `json:"gas_anl_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Fluid_type                 *string  `json:"fluid_type" default:""`
	Gas_gravity                *float64 `json:"gas_gravity" default:""`
	Gas_gravity_ouom           *string  `json:"gas_gravity_ouom" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Preferred_ind              *string  `json:"preferred_ind" default:""`
	Problem_ind                *string  `json:"problem_ind" default:""`
	Pseudo_critical_press      *float64 `json:"pseudo_critical_press" default:""`
	Pseudo_critical_press_ouom *string  `json:"pseudo_critical_press_ouom" default:""`
	Pseudo_critical_temp       *float64 `json:"pseudo_critical_temp" default:""`
	Pseudo_critical_temp_ouom  *string  `json:"pseudo_critical_temp_ouom" default:""`
	Remark                     *string  `json:"remark" default:""`
	Source                     *string  `json:"source" default:""`
	Step_seq_no                *int     `json:"step_seq_no" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
