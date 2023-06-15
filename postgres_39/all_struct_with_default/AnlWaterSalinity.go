package dto

type Anl_water_salinity struct {
	Analysis_id           string   `json:"analysis_id" default:"source"`
	Anl_source            string   `json:"anl_source" default:"source"`
	Water_salinity_obs_no int      `json:"water_salinity_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Calculated_ind        *string  `json:"calculated_ind" default:""`
	Calculate_method_id   *string  `json:"calculate_method_id" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ignition_ind          *string  `json:"ignition_ind" default:""`
	Measured_ind          *string  `json:"measured_ind" default:""`
	Measured_temp         *float64 `json:"measured_temp" default:""`
	Measured_temp_ouom    *string  `json:"measured_temp_ouom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Preferred_ind         *string  `json:"preferred_ind" default:""`
	Problem_ind           *string  `json:"problem_ind" default:""`
	Remark                *string  `json:"remark" default:""`
	Reported_ind          *string  `json:"reported_ind" default:""`
	Salinity              *float64 `json:"salinity" default:""`
	Salinity_ouom         *string  `json:"salinity_ouom" default:""`
	Salinity_type         *string  `json:"salinity_type" default:""`
	Salinity_uom          *string  `json:"salinity_uom" default:""`
	Source                *string  `json:"source" default:""`
	Step_seq_no           *int     `json:"step_seq_no" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
