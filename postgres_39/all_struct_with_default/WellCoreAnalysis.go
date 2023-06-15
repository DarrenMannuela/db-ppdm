package dto

type Well_core_analysis struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Source                string   `json:"source" default:"source"`
	Core_id               string   `json:"core_id" default:"source"`
	Analysis_obs_no       int      `json:"analysis_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Analysis_date         *string  `json:"analysis_date" default:""`
	Analyzing_company     *string  `json:"analyzing_company" default:""`
	Analyzing_company_loc *string  `json:"analyzing_company_loc" default:""`
	Analyzing_file_num    *string  `json:"analyzing_file_num" default:""`
	Core_analyst_name     *string  `json:"core_analyst_name" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Primary_sample_type   *string  `json:"primary_sample_type" default:""`
	Remark                *string  `json:"remark" default:""`
	Sample_diameter       *float64 `json:"sample_diameter" default:""`
	Sample_diameter_ouom  *string  `json:"sample_diameter_ouom" default:""`
	Sample_length         *float64 `json:"sample_length" default:""`
	Sample_length_ouom    *string  `json:"sample_length_ouom" default:""`
	Sample_shape          *string  `json:"sample_shape" default:""`
	Second_sample_type    *string  `json:"second_sample_type" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
