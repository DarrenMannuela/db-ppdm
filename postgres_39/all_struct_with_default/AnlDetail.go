package dto

type Anl_detail struct {
	Analysis_id          string   `json:"analysis_id" default:"source"`
	Anl_source           string   `json:"anl_source" default:"source"`
	Detail_obs_no        int      `json:"detail_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Anl_detail_type      *string  `json:"anl_detail_type" default:""`
	Average_ratio_value  *float64 `json:"average_ratio_value" default:""`
	Average_value        *float64 `json:"average_value" default:""`
	Average_value_ouom   *string  `json:"average_value_ouom" default:""`
	Average_value_uom    *string  `json:"average_value_uom" default:""`
	Calculated_ind       *string  `json:"calculated_ind" default:""`
	Calculate_method_id  *string  `json:"calculate_method_id" default:""`
	Date_format_desc     *string  `json:"date_format_desc" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Max_date             *string  `json:"max_date" default:""`
	Max_ratio_value      *float64 `json:"max_ratio_value" default:""`
	Max_value            *float64 `json:"max_value" default:""`
	Max_value_ouom       *string  `json:"max_value_ouom" default:""`
	Max_value_uom        *string  `json:"max_value_uom" default:""`
	Measurement_type     *string  `json:"measurement_type" default:""`
	Min_date             *string  `json:"min_date" default:""`
	Min_ratio_value      *float64 `json:"min_ratio_value" default:""`
	Min_value            *float64 `json:"min_value" default:""`
	Min_value_ouom       *string  `json:"min_value_ouom" default:""`
	Min_value_uom        *string  `json:"min_value_uom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Problem_ind          *string  `json:"problem_ind" default:""`
	Reference_value      *float64 `json:"reference_value" default:""`
	Reference_value_ouom *string  `json:"reference_value_ouom" default:""`
	Reference_value_type *string  `json:"reference_value_type" default:""`
	Reference_value_uom  *string  `json:"reference_value_uom" default:""`
	Remark               *string  `json:"remark" default:""`
	Reported_ind         *string  `json:"reported_ind" default:""`
	Source               *string  `json:"source" default:""`
	Step_seq_no          *int     `json:"step_seq_no" default:""`
	Substance_id         *string  `json:"substance_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
