package dto

type Anl_isotope struct {
	Analysis_id         string   `json:"analysis_id" default:"source"`
	Anl_source          string   `json:"anl_source" default:"source"`
	Isotope_obs_no      int      `json:"isotope_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Calculate_method_id *string  `json:"calculate_method_id" default:""`
	Delta_notation_ind  *string  `json:"delta_notation_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Isotope_value       *float64 `json:"isotope_value" default:""`
	Isotope_value_ouom  *string  `json:"isotope_value_ouom" default:""`
	Isotope_value_uom   *string  `json:"isotope_value_uom" default:""`
	Measurement_type    *string  `json:"measurement_type" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Preferred_ind       *string  `json:"preferred_ind" default:""`
	Problem_ind         *string  `json:"problem_ind" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Standard_id         *string  `json:"standard_id" default:""`
	Step_seq_no         *int     `json:"step_seq_no" default:""`
	Substance_id        *string  `json:"substance_id" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
