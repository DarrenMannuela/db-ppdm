package dto

import (
	"time"
)

type Well_log_curve struct {
	Uwi                    string     `json:"uwi"`
	Curve_id               string     `json:"curve_id"`
	Acquired_for_ba_id     *string    `json:"acquired_for_ba_id"`
	Active_ind             *string    `json:"active_ind"`
	Api_code_system        *string    `json:"api_code_system"`
	Api_curve_class        *string    `json:"api_curve_class"`
	Api_curve_code         *int       `json:"api_curve_code"`
	Api_curve_modifier     *float64   `json:"api_curve_modifier"`
	Api_log_code           *int       `json:"api_log_code"`
	Base_curve_ind         *string    `json:"base_curve_ind"`
	Bypass_ind             *string    `json:"bypass_ind"`
	Cased_hole_ind         *string    `json:"cased_hole_ind"`
	Composite_ind          *string    `json:"composite_ind"`
	Curve_ouom             *string    `json:"curve_ouom"`
	Curve_output_type      *string    `json:"curve_output_type"`
	Curve_quality          *string    `json:"curve_quality"`
	Date_format_desc       *time.Time `json:"date_format_desc"`
	Dictionary_id          *string    `json:"dictionary_id"`
	Dict_curve_id          *string    `json:"dict_curve_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Explicit_index_ind     *string    `json:"explicit_index_ind"`
	First_good_value       *float64   `json:"first_good_value"`
	First_good_value_index *float64   `json:"first_good_value_index"`
	Frame_id               *string    `json:"frame_id"`
	Good_value_type        *string    `json:"good_value_type"`
	Index_curve_id         *string    `json:"index_curve_id"`
	Index_ouom             *string    `json:"index_ouom"`
	Index_uom              *string    `json:"index_uom"`
	Job_id                 *string    `json:"job_id"`
	Last_good_value        *float64   `json:"last_good_value"`
	Last_good_value_index  *float64   `json:"last_good_value_index"`
	Log_tool_pass_no       *int       `json:"log_tool_pass_no"`
	Log_tool_type          *string    `json:"log_tool_type"`
	Max_index              *float64   `json:"max_index"`
	Max_value              *float64   `json:"max_value"`
	Max_value_index        *float64   `json:"max_value_index"`
	Mean_value             *float64   `json:"mean_value"`
	Mean_value_std_dev     *float64   `json:"mean_value_std_dev"`
	Min_index              *float64   `json:"min_index"`
	Min_value              *float64   `json:"min_value"`
	Min_value_index        *float64   `json:"min_value_index"`
	Multiple_index_ind     *string    `json:"multiple_index_ind"`
	Mwd_ind                *string    `json:"mwd_ind"`
	Null_count             *int       `json:"null_count"`
	Null_representation    *string    `json:"null_representation"`
	Order_seq_no           *int       `json:"order_seq_no"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Primary_index_type     *string    `json:"primary_index_type"`
	Remark                 *string    `json:"remark"`
	Reported_desc          *string    `json:"reported_desc"`
	Reported_mnemonic      *string    `json:"reported_mnemonic"`
	Reported_unit_mnemonic *string    `json:"reported_unit_mnemonic"`
	Source                 *string    `json:"source"`
	Trip_obs_no            *int       `json:"trip_obs_no"`
	Value_count            *int       `json:"value_count"`
	Well_log_id            *string    `json:"well_log_id"`
	Well_log_job_source    *string    `json:"well_log_job_source"`
	Well_log_source        *string    `json:"well_log_source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
