package dto

type Anl_oil_fraction struct {
	Analysis_id                 string   `json:"analysis_id" default:"source"`
	Anl_source                  string   `json:"anl_source" default:"source"`
	Distill_fraction_obs_no     int      `json:"distill_fraction_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Calculated_ind              *string  `json:"calculated_ind" default:""`
	Calculate_method_id         *string  `json:"calculate_method_id" default:""`
	Distill_temp                *float64 `json:"distill_temp" default:""`
	Distill_temp_ouom           *string  `json:"distill_temp_ouom" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Problem_ind                 *string  `json:"problem_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Reported_ind                *string  `json:"reported_ind" default:""`
	Source                      *string  `json:"source" default:""`
	Step_seq_no                 *int     `json:"step_seq_no" default:""`
	Substance_id                *string  `json:"substance_id" default:""`
	Vol_fraction_distilled      *float64 `json:"vol_fraction_distilled" default:""`
	Vol_fraction_distilled_ouom *string  `json:"vol_fraction_distilled_ouom" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
