package dto

type Anl_gas_composition struct {
	Analysis_id                string   `json:"analysis_id" default:"source"`
	Anl_source                 string   `json:"anl_source" default:"source"`
	Gas_anl_composition_obs_no int      `json:"gas_anl_composition_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Calculated_ind             *string  `json:"calculated_ind" default:""`
	Calculate_method_id        *string  `json:"calculate_method_id" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Not_present_ind            *string  `json:"not_present_ind" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Problem_ind                *string  `json:"problem_ind" default:""`
	Remark                     *string  `json:"remark" default:""`
	Reported_ind               *string  `json:"reported_ind" default:""`
	Source                     *string  `json:"source" default:""`
	Step_seq_no                *int     `json:"step_seq_no" default:""`
	Substance_amount           *float64 `json:"substance_amount" default:""`
	Substance_amount_ouom      *string  `json:"substance_amount_ouom" default:""`
	Substance_amount_uom       *string  `json:"substance_amount_uom" default:""`
	Substance_id               *string  `json:"substance_id" default:""`
	Substance_percent          *float64 `json:"substance_percent" default:""`
	Trace_ind                  *string  `json:"trace_ind" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
