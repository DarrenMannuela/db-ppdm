package dto

import (
	"time"
)

type Anl_oil_fraction struct {
	Analysis_id                 string     `json:"analysis_id"`
	Anl_source                  string     `json:"anl_source"`
	Distill_fraction_obs_no     int        `json:"distill_fraction_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Calculated_ind              *string    `json:"calculated_ind"`
	Calculate_method_id         *string    `json:"calculate_method_id"`
	Distill_temp                *float64   `json:"distill_temp"`
	Distill_temp_ouom           *string    `json:"distill_temp_ouom"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Problem_ind                 *string    `json:"problem_ind"`
	Remark                      *string    `json:"remark"`
	Reported_ind                *string    `json:"reported_ind"`
	Source                      *string    `json:"source"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	Substance_id                *string    `json:"substance_id"`
	Vol_fraction_distilled      *float64   `json:"vol_fraction_distilled"`
	Vol_fraction_distilled_ouom *string    `json:"vol_fraction_distilled_ouom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
