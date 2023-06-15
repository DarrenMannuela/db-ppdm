package dto

import (
	"time"
)

type Anl_isotope struct {
	Analysis_id         string     `json:"analysis_id"`
	Anl_source          string     `json:"anl_source"`
	Isotope_obs_no      int        `json:"isotope_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Calculate_method_id *string    `json:"calculate_method_id"`
	Delta_notation_ind  *string    `json:"delta_notation_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Isotope_value       *float64   `json:"isotope_value"`
	Isotope_value_ouom  *string    `json:"isotope_value_ouom"`
	Isotope_value_uom   *string    `json:"isotope_value_uom"`
	Measurement_type    *string    `json:"measurement_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Preferred_ind       *string    `json:"preferred_ind"`
	Problem_ind         *string    `json:"problem_ind"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Standard_id         *string    `json:"standard_id"`
	Step_seq_no         *int       `json:"step_seq_no"`
	Substance_id        *string    `json:"substance_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
