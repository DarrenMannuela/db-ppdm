package dto

type Well_log_curve_value struct {
	Uwi                string   `json:"uwi" default:"source"`
	Curve_id           string   `json:"curve_id" default:"source"`
	Sample_id          string   `json:"sample_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Index_type         *string  `json:"index_type" default:""`
	Index_value        *float64 `json:"index_value" default:""`
	Index_value_uom    *string  `json:"index_value_uom" default:""`
	Measured_value     *float64 `json:"measured_value" default:""`
	Measured_value_uom *string  `json:"measured_value_uom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
