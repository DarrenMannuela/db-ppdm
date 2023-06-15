package dto

import (
	"time"
)

type Anl_paleo_maturity struct {
	Analysis_id                string     `json:"analysis_id"`
	Anl_source                 string     `json:"anl_source"`
	Maturation_obs_no          int        `json:"maturation_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Fluor_color                *string    `json:"fluor_color"`
	Fluor_intensity_desc       *string    `json:"fluor_intensity_desc"`
	Fluor_intensity_value      *float64   `json:"fluor_intensity_value"`
	Fluor_intensity_value_ouom *string    `json:"fluor_intensity_value_ouom"`
	Fluor_intensity_value_uom  *string    `json:"fluor_intensity_value_uom"`
	Maturity_method_type       *string    `json:"maturity_method_type"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_ind              *string    `json:"preferred_ind"`
	Problem_ind                *string    `json:"problem_ind"`
	Remark                     *string    `json:"remark"`
	Sci_color                  *string    `json:"sci_color"`
	Sci_max                    *float64   `json:"sci_max"`
	Sci_max_ouom               *string    `json:"sci_max_ouom"`
	Sci_max_uom                *string    `json:"sci_max_uom"`
	Sci_min                    *float64   `json:"sci_min"`
	Sci_min_ouom               *string    `json:"sci_min_ouom"`
	Sci_min_uom                *string    `json:"sci_min_uom"`
	Source                     *string    `json:"source"`
	Step_seq_no                *int       `json:"step_seq_no"`
	Substance_id               *string    `json:"substance_id"`
	Tai_color                  *string    `json:"tai_color"`
	Tai_max                    *float64   `json:"tai_max"`
	Tai_max_ouom               *string    `json:"tai_max_ouom"`
	Tai_max_uom                *string    `json:"tai_max_uom"`
	Tai_min                    *float64   `json:"tai_min"`
	Tai_min_ouom               *string    `json:"tai_min_ouom"`
	Tai_min_uom                *string    `json:"tai_min_uom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
