package dto

import (
	"time"
)

type Anl_elemental_detail struct {
	Analysis_id         string     `json:"analysis_id"`
	Anl_source          string     `json:"anl_source"`
	Detail_anl_obs_no   int        `json:"detail_anl_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Analysis_value      *float64   `json:"analysis_value"`
	Analysis_value_ouom *string    `json:"analysis_value_ouom"`
	Analysis_value_type *string    `json:"analysis_value_type"`
	Analysis_value_uom  *string    `json:"analysis_value_uom"`
	Average_value       *float64   `json:"average_value"`
	Average_value_ouom  *string    `json:"average_value_ouom"`
	Average_value_uom   *string    `json:"average_value_uom"`
	Calculated_ind      *string    `json:"calculated_ind"`
	Calculate_method_id *string    `json:"calculate_method_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Max_date            *time.Time `json:"max_date"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
	Measurement_type    *string    `json:"measurement_type"`
	Min_date            *time.Time `json:"min_date"`
	Min_value           *float64   `json:"min_value"`
	Min_value_ouom      *string    `json:"min_value_ouom"`
	Min_value_uom       *string    `json:"min_value_uom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Preferred_ind       *string    `json:"preferred_ind"`
	Remark              *string    `json:"remark"`
	Reported_ind        *string    `json:"reported_ind"`
	Source              *string    `json:"source"`
	Step_seq_no         *int       `json:"step_seq_no"`
	Substance_id        *string    `json:"substance_id"`
	Valid_value_ind     *string    `json:"valid_value_ind"`
	Valid_value_remark  *string    `json:"valid_value_remark"`
	Value_code          *string    `json:"value_code"`
	Value_desc          *string    `json:"value_desc"`
	Value_quality       *string    `json:"value_quality"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
