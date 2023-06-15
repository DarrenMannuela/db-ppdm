package dto

import (
	"time"
)

type Anl_liquid_chro struct {
	Analysis_id               string     `json:"analysis_id"`
	Anl_source                string     `json:"anl_source"`
	Chro_obs_no               int        `json:"chro_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Aged_oil_used             *float64   `json:"aged_oil_used"`
	Aged_oil_used_uom         *string    `json:"aged_oil_used_uom"`
	Aged_weight_of_oil        *float64   `json:"aged_weight_of_oil"`
	Aged_weight_of_oil_uom    *string    `json:"aged_weight_of_oil_uom"`
	Aromatic_hc_value         *float64   `json:"aromatic_hc_value"`
	Aromatic_hc_value_uom     *string    `json:"aromatic_hc_value_uom"`
	Asphaltene_free_value     *float64   `json:"asphaltene_free_value"`
	Asphaltene_free_value_uom *string    `json:"asphaltene_free_value_uom"`
	Asphaltene_value          *float64   `json:"asphaltene_value"`
	Asphaltene_value_uom      *string    `json:"asphaltene_value_uom"`
	Effective_date            *time.Time `json:"effective_date"`
	Eom_used                  *float64   `json:"eom_used"`
	Eom_used_uom              *string    `json:"eom_used_uom"`
	Eom_value                 *float64   `json:"eom_value"`
	Eom_value_uom             *string    `json:"eom_value_uom"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Loss_on_column            *float64   `json:"loss_on_column"`
	Loss_on_column_uom        *string    `json:"loss_on_column_uom"`
	Measurement_type          *string    `json:"measurement_type"`
	Naphthene_value           *float64   `json:"naphthene_value"`
	Naphthene_value_uom       *string    `json:"naphthene_value_uom"`
	Nso_value                 *float64   `json:"nso_value"`
	Nso_value_uom             *string    `json:"nso_value_uom"`
	N_alkane_fraction         *float64   `json:"n_alkane_fraction"`
	N_alkane_fraction_uom     *string    `json:"n_alkane_fraction_uom"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Preferred_ind             *string    `json:"preferred_ind"`
	Problem_ind               *string    `json:"problem_ind"`
	Refraction_factor         *float64   `json:"refraction_factor"`
	Remark                    *string    `json:"remark"`
	Reported_hc_fraction      *float64   `json:"reported_hc_fraction"`
	Reported_hc_fraction_uom  *string    `json:"reported_hc_fraction_uom"`
	Reported_hc_in_toc        *float64   `json:"reported_hc_in_toc"`
	Reported_hc_in_toc_uom    *string    `json:"reported_hc_in_toc_uom"`
	Saturated_hc_value        *float64   `json:"saturated_hc_value"`
	Saturated_hc_value_uom    *string    `json:"saturated_hc_value_uom"`
	Source                    *string    `json:"source"`
	Step_seq_no               *int       `json:"step_seq_no"`
	Total_soluble_extract     *float64   `json:"total_soluble_extract"`
	Total_soluble_extract_uom *string    `json:"total_soluble_extract_uom"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
