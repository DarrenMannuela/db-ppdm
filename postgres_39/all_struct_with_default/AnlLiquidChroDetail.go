package dto

type Anl_liquid_chro_detail struct {
	Analysis_id               string   `json:"analysis_id" default:"source"`
	Anl_source                string   `json:"anl_source" default:"source"`
	Detail_obs_no             int      `json:"detail_obs_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Anl_value                 *float64 `json:"anl_value" default:""`
	Anl_value_ouom            *string  `json:"anl_value_ouom" default:""`
	Anl_value_remark          *string  `json:"anl_value_remark" default:""`
	Anl_value_uom             *string  `json:"anl_value_uom" default:""`
	Calculated_ind            *string  `json:"calculated_ind" default:""`
	Calculate_method_id       *string  `json:"calculate_method_id" default:""`
	Chro_property_type        *string  `json:"chro_property_type" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Error_bar_value           *float64 `json:"error_bar_value" default:""`
	Error_bar_value_ouom      *string  `json:"error_bar_value_ouom" default:""`
	Error_bar_value_uom       *string  `json:"error_bar_value_uom" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Measured_ind              *string  `json:"measured_ind" default:""`
	Measurement_type          *string  `json:"measurement_type" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Preferred_ind             *string  `json:"preferred_ind" default:""`
	Problem_ind               *string  `json:"problem_ind" default:""`
	Quantif_additive_amt      *float64 `json:"quantif_additive_amt" default:""`
	Quantif_additive_amt_ouom *string  `json:"quantif_additive_amt_ouom" default:""`
	Quantif_additive_amt_uom  *string  `json:"quantif_additive_amt_uom" default:""`
	Quantif_additive_desc     *string  `json:"quantif_additive_desc" default:""`
	Quantif_additive_type     *string  `json:"quantif_additive_type" default:""`
	Remark                    *string  `json:"remark" default:""`
	Reported_ind              *string  `json:"reported_ind" default:""`
	Source                    *string  `json:"source" default:""`
	Step_seq_no               *int     `json:"step_seq_no" default:""`
	Substance_id              *string  `json:"substance_id" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
