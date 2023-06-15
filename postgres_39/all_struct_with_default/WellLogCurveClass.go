package dto

type Well_log_curve_class struct {
	Curve_class_id     string  `json:"curve_class_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Class_mnemonic     *string `json:"class_mnemonic" default:""`
	Curve_class_name   *string `json:"curve_class_name" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Quantity_id        *string `json:"quantity_id" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	System_type        *string `json:"system_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
