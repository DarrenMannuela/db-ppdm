package dto

import (
	"time"
)

type Sample_lith_desc struct {
	Sample_id                string     `json:"sample_id"`
	Description_id           string     `json:"description_id"`
	Active_ind               *string    `json:"active_ind"`
	Color_intensity          *string    `json:"color_intensity"`
	Color_type               *string    `json:"color_type"`
	Contaminant              *string    `json:"contaminant"`
	Depositional_environment *string    `json:"depositional_environment"`
	Diagenesis_type          *string    `json:"diagenesis_type"`
	Ecozone_id               *string    `json:"ecozone_id"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Lithology                *string    `json:"lithology"`
	Oil_stain                *string    `json:"oil_stain"`
	Percent                  *float64   `json:"percent"`
	Permeability_type        *string    `json:"permeability_type"`
	Porosity_type            *string    `json:"porosity_type"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Rock_matrix              *string    `json:"rock_matrix"`
	Rock_type                *string    `json:"rock_type"`
	Source                   *string    `json:"source"`
	Strat_name_set_id        *string    `json:"strat_name_set_id"`
	Strat_unit_id            *string    `json:"strat_unit_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
