package dto

import (
	"time"
)

type Anl_water_detail struct {
	Analysis_id          string     `json:"analysis_id"`
	Anl_source           string     `json:"anl_source"`
	Detail_obs_no        int        `json:"detail_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Analysis_press       *float64   `json:"analysis_press"`
	Analysis_press_ouom  *string    `json:"analysis_press_ouom"`
	Analysis_press_uom   *string    `json:"analysis_press_uom"`
	Analysis_temp        *float64   `json:"analysis_temp"`
	Analysis_temp_ouom   *string    `json:"analysis_temp_ouom"`
	Analysis_temp_uom    *string    `json:"analysis_temp_uom"`
	Analysis_value       *float64   `json:"analysis_value"`
	Analysis_value_ouom  *string    `json:"analysis_value_ouom"`
	Analysis_value_type  *string    `json:"analysis_value_type"`
	Analysis_value_uom   *string    `json:"analysis_value_uom"`
	Anl_value_remark     *string    `json:"anl_value_remark"`
	Calculated_value_ind *string    `json:"calculated_value_ind"`
	Calculate_method_id  *string    `json:"calculate_method_id"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Measured_value_ind   *string    `json:"measured_value_ind"`
	Measurement_type     *string    `json:"measurement_type"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Problem_ind          *string    `json:"problem_ind"`
	Remark               *string    `json:"remark"`
	Reported_value_ind   *string    `json:"reported_value_ind"`
	Source               *string    `json:"source"`
	Step_seq_no          *int       `json:"step_seq_no"`
	Substance_id         *string    `json:"substance_id"`
	Water_property       *string    `json:"water_property"`
	Water_property_code  *string    `json:"water_property_code"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
