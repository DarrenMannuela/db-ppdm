package dto

import (
	"time"
)

type Anl_detail struct {
	Analysis_id          string     `json:"analysis_id"`
	Anl_source           string     `json:"anl_source"`
	Detail_obs_no        int        `json:"detail_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Anl_detail_type      *string    `json:"anl_detail_type"`
	Average_ratio_value  *float64   `json:"average_ratio_value"`
	Average_value        *float64   `json:"average_value"`
	Average_value_ouom   *string    `json:"average_value_ouom"`
	Average_value_uom    *string    `json:"average_value_uom"`
	Calculated_ind       *string    `json:"calculated_ind"`
	Calculate_method_id  *string    `json:"calculate_method_id"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Max_date             *time.Time `json:"max_date"`
	Max_ratio_value      *float64   `json:"max_ratio_value"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Measurement_type     *string    `json:"measurement_type"`
	Min_date             *time.Time `json:"min_date"`
	Min_ratio_value      *float64   `json:"min_ratio_value"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Problem_ind          *string    `json:"problem_ind"`
	Reference_value      *float64   `json:"reference_value"`
	Reference_value_ouom *string    `json:"reference_value_ouom"`
	Reference_value_type *string    `json:"reference_value_type"`
	Reference_value_uom  *string    `json:"reference_value_uom"`
	Remark               *string    `json:"remark"`
	Reported_ind         *string    `json:"reported_ind"`
	Source               *string    `json:"source"`
	Step_seq_no          *int       `json:"step_seq_no"`
	Substance_id         *string    `json:"substance_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
