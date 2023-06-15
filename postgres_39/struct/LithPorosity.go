package dto

import (
	"time"
)

type Lith_porosity struct {
	Lithology_log_id     string     `json:"lithology_log_id"`
	Lith_log_source      string     `json:"lith_log_source"`
	Depth_obs_no         int        `json:"depth_obs_no"`
	Rock_type            string     `json:"rock_type"`
	Rock_type_obs_no     int        `json:"rock_type_obs_no"`
	Porosity_grade       float64    `json:"porosity_grade"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Matrix_class         *string    `json:"matrix_class"`
	Percent_of_sample    *float64   `json:"percent_of_sample"`
	Porosity_scale       *string    `json:"porosity_scale"`
	Porosity_type        *string    `json:"porosity_type"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Sample_rel_abundance *string    `json:"sample_rel_abundance"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
