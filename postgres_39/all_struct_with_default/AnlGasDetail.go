package dto

type Anl_gas_detail struct {
	Analysis_id           string   `json:"analysis_id" default:"source"`
	Anl_source            string   `json:"anl_source" default:"source"`
	Gas_anl_detail_obs_no int      `json:"gas_anl_detail_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Analysis_property     *string  `json:"analysis_property" default:""`
	Analysis_value        *float64 `json:"analysis_value" default:""`
	Analysis_value_code   *string  `json:"analysis_value_code" default:""`
	Analysis_value_ouom   *string  `json:"analysis_value_ouom" default:""`
	Analysis_value_uom    *string  `json:"analysis_value_uom" default:""`
	Anl_value_remark      *string  `json:"anl_value_remark" default:""`
	Calculate_method_id   *string  `json:"calculate_method_id" default:""`
	Date_format_desc      *string  `json:"date_format_desc" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Max_value             *float64 `json:"max_value" default:""`
	Max_value_ouom        *string  `json:"max_value_ouom" default:""`
	Max_value_uom         *string  `json:"max_value_uom" default:""`
	Measurement_type      *string  `json:"measurement_type" default:""`
	Min_value             *float64 `json:"min_value" default:""`
	Min_value_ouom        *string  `json:"min_value_ouom" default:""`
	Min_value_uom         *string  `json:"min_value_uom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Step_seq_no           *int     `json:"step_seq_no" default:""`
	Substance_id          *string  `json:"substance_id" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
