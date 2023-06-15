package dto

import (
	"time"
)

type Anl_gas_chro struct {
	Analysis_id               string     `json:"analysis_id"`
	Anl_source                string     `json:"anl_source"`
	Anl_obs_no                int        `json:"anl_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Anl_value                 *float64   `json:"anl_value"`
	Anl_value_code            *string    `json:"anl_value_code"`
	Anl_value_ouom            *string    `json:"anl_value_ouom"`
	Anl_value_remark          *string    `json:"anl_value_remark"`
	Anl_value_type            *string    `json:"anl_value_type"`
	Anl_value_uom             *string    `json:"anl_value_uom"`
	Calculated_ind            *string    `json:"calculated_ind"`
	Calculate_method_id       *string    `json:"calculate_method_id"`
	Daughter_ion_id           *string    `json:"daughter_ion_id"`
	Effective_date            *time.Time `json:"effective_date"`
	Error_bar_value           *float64   `json:"error_bar_value"`
	Error_bar_value_ouom      *string    `json:"error_bar_value_ouom"`
	Error_bar_value_uom       *string    `json:"error_bar_value_uom"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Gas_anl_value_type        *string    `json:"gas_anl_value_type"`
	Measurement_type          *string    `json:"measurement_type"`
	Parent_ion_id             *string    `json:"parent_ion_id"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Preferred_ind             *string    `json:"preferred_ind"`
	Problem_ind               *string    `json:"problem_ind"`
	Qualifier_ion_id          *string    `json:"qualifier_ion_id"`
	Quantif_additive_amt      *float64   `json:"quantif_additive_amt"`
	Quantif_additive_amt_ouom *string    `json:"quantif_additive_amt_ouom"`
	Quantif_additive_amt_uom  *string    `json:"quantif_additive_amt_uom"`
	Quantif_additive_desc     *string    `json:"quantif_additive_desc"`
	Quantif_additive_type     *string    `json:"quantif_additive_type"`
	Remark                    *string    `json:"remark"`
	Reported_ind              *string    `json:"reported_ind"`
	Source                    *string    `json:"source"`
	Step_seq_no               *int       `json:"step_seq_no"`
	Substance_id              *string    `json:"substance_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
