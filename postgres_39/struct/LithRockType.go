package dto

import (
	"time"
)

type Lith_rock_type struct {
	Lithology_log_id          string     `json:"lithology_log_id"`
	Lith_log_source           string     `json:"lith_log_source"`
	Depth_obs_no              int        `json:"depth_obs_no"`
	Rock_type                 string     `json:"rock_type"`
	Rock_type_obs_no          int        `json:"rock_type_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Class                     *string    `json:"class"`
	Class_modifier            *string    `json:"class_modifier"`
	Consolidation             *string    `json:"consolidation"`
	Cut_color                 *string    `json:"cut_color"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Fluorescence_color        *string    `json:"fluorescence_color"`
	Fluorescence_distribution *string    `json:"fluorescence_distribution"`
	Fluorescence_rate         *string    `json:"fluorescence_rate"`
	Framework_percent         *float64   `json:"framework_percent"`
	Matrix_percent            *float64   `json:"matrix_percent"`
	Oil_stain                 *string    `json:"oil_stain"`
	Permeability_quality      *string    `json:"permeability_quality"`
	Porosity_grade_percent    *float64   `json:"porosity_grade_percent"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Prim_porosity_type        *string    `json:"prim_porosity_type"`
	Remark                    *string    `json:"remark"`
	Residue_color             *string    `json:"residue_color"`
	Residue_percent           *float64   `json:"residue_percent"`
	Rock_class_scheme         *string    `json:"rock_class_scheme"`
	Rock_matrix               *string    `json:"rock_matrix"`
	Rock_profile              *string    `json:"rock_profile"`
	Rock_rel_abundance        *string    `json:"rock_rel_abundance"`
	Rock_type_percent         *float64   `json:"rock_type_percent"`
	Rounding                  *string    `json:"rounding"`
	Sec_porosity_type         *string    `json:"sec_porosity_type"`
	Solid_hcarbon_percent     *float64   `json:"solid_hcarbon_percent"`
	Solid_hcarbon_type        *string    `json:"solid_hcarbon_type"`
	Sorting                   *string    `json:"sorting"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
