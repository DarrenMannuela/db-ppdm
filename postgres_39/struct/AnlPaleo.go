package dto

import (
	"time"
)

type Anl_paleo struct {
	Analysis_id            string     `json:"analysis_id"`
	Anl_source             string     `json:"anl_source"`
	Palynology_obs_no      int        `json:"palynology_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Amount_type            *string    `json:"amount_type"`
	Calculated_ind         *string    `json:"calculated_ind"`
	Calculate_method_id    *string    `json:"calculate_method_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Problem_ind            *string    `json:"problem_ind"`
	Remark                 *string    `json:"remark"`
	Reported_ind           *string    `json:"reported_ind"`
	Sample_total_value     *float64   `json:"sample_total_value"`
	Sample_total_value_uom *string    `json:"sample_total_value_uom"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Substance_id           *string    `json:"substance_id"`
	Total_value            *float64   `json:"total_value"`
	Total_value_ouom       *string    `json:"total_value_ouom"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
