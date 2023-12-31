package dto

type Well_log_curve struct {
	Uwi                    string   `json:"uwi" default:"source"`
	Curve_id               string   `json:"curve_id" default:"source"`
	Acquired_for_ba_id     *string  `json:"acquired_for_ba_id" default:""`
	Active_ind             *string  `json:"active_ind" default:""`
	Api_code_system        *string  `json:"api_code_system" default:""`
	Api_curve_class        *string  `json:"api_curve_class" default:""`
	Api_curve_code         *int     `json:"api_curve_code" default:""`
	Api_curve_modifier     *float64 `json:"api_curve_modifier" default:""`
	Api_log_code           *int     `json:"api_log_code" default:""`
	Base_curve_ind         *string  `json:"base_curve_ind" default:""`
	Bypass_ind             *string  `json:"bypass_ind" default:""`
	Cased_hole_ind         *string  `json:"cased_hole_ind" default:""`
	Composite_ind          *string  `json:"composite_ind" default:""`
	Curve_ouom             *string  `json:"curve_ouom" default:""`
	Curve_output_type      *string  `json:"curve_output_type" default:""`
	Curve_quality          *string  `json:"curve_quality" default:""`
	Date_format_desc       *string  `json:"date_format_desc" default:""`
	Dictionary_id          *string  `json:"dictionary_id" default:""`
	Dict_curve_id          *string  `json:"dict_curve_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Explicit_index_ind     *string  `json:"explicit_index_ind" default:""`
	First_good_value       *float64 `json:"first_good_value" default:""`
	First_good_value_index *float64 `json:"first_good_value_index" default:""`
	Frame_id               *string  `json:"frame_id" default:""`
	Good_value_type        *string  `json:"good_value_type" default:""`
	Index_curve_id         *string  `json:"index_curve_id" default:""`
	Index_ouom             *string  `json:"index_ouom" default:""`
	Index_uom              *string  `json:"index_uom" default:""`
	Job_id                 *string  `json:"job_id" default:""`
	Last_good_value        *float64 `json:"last_good_value" default:""`
	Last_good_value_index  *float64 `json:"last_good_value_index" default:""`
	Log_tool_pass_no       *int     `json:"log_tool_pass_no" default:""`
	Log_tool_type          *string  `json:"log_tool_type" default:""`
	Max_index              *float64 `json:"max_index" default:""`
	Max_value              *float64 `json:"max_value" default:""`
	Max_value_index        *float64 `json:"max_value_index" default:""`
	Mean_value             *float64 `json:"mean_value" default:""`
	Mean_value_std_dev     *float64 `json:"mean_value_std_dev" default:""`
	Min_index              *float64 `json:"min_index" default:""`
	Min_value              *float64 `json:"min_value" default:""`
	Min_value_index        *float64 `json:"min_value_index" default:""`
	Multiple_index_ind     *string  `json:"multiple_index_ind" default:""`
	Mwd_ind                *string  `json:"mwd_ind" default:""`
	Null_count             *int     `json:"null_count" default:""`
	Null_representation    *string  `json:"null_representation" default:""`
	Order_seq_no           *int     `json:"order_seq_no" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Preferred_ind          *string  `json:"preferred_ind" default:""`
	Primary_index_type     *string  `json:"primary_index_type" default:""`
	Remark                 *string  `json:"remark" default:""`
	Reported_desc          *string  `json:"reported_desc" default:""`
	Reported_mnemonic      *string  `json:"reported_mnemonic" default:""`
	Reported_unit_mnemonic *string  `json:"reported_unit_mnemonic" default:""`
	Source                 *string  `json:"source" default:""`
	Trip_obs_no            *int     `json:"trip_obs_no" default:""`
	Value_count            *int     `json:"value_count" default:""`
	Well_log_id            *string  `json:"well_log_id" default:""`
	Well_log_job_source    *string  `json:"well_log_job_source" default:""`
	Well_log_source        *string  `json:"well_log_source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
