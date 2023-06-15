package dto

import (
	"time"
)

type Anl_maceral_maturity struct {
	Analysis_id                 string     `json:"analysis_id"`
	Anl_source                  string     `json:"anl_source"`
	Measurement_obs_no          int        `json:"measurement_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Coal_rank_id                *string    `json:"coal_rank_id"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Fluor_color                 *string    `json:"fluor_color"`
	Fluor_intensity_desc        *string    `json:"fluor_intensity_desc"`
	Fluor_intensity_value       *float64   `json:"fluor_intensity_value"`
	Fluor_intensity_value_ouom  *string    `json:"fluor_intensity_value_ouom"`
	Mean_max_reflectance        *float64   `json:"mean_max_reflectance"`
	Mean_random_reflectance     *float64   `json:"mean_random_reflectance"`
	Measurements_count          *int       `json:"measurements_count"`
	Population_name             *string    `json:"population_name"`
	Population_num              *string    `json:"population_num"`
	Population_obs_no           *int       `json:"population_obs_no"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Preferred_ind               *string    `json:"preferred_ind"`
	Problem_ind                 *string    `json:"problem_ind"`
	Reflect_max_meas_value      *float64   `json:"reflect_max_meas_value"`
	Reflect_max_meas_value_ouom *string    `json:"reflect_max_meas_value_ouom"`
	Reflect_max_meas_value_uom  *string    `json:"reflect_max_meas_value_uom"`
	Reflect_meas_value          *float64   `json:"reflect_meas_value"`
	Reflect_meas_value_ouom     *string    `json:"reflect_meas_value_ouom"`
	Reflect_meas_value_uom      *string    `json:"reflect_meas_value_uom"`
	Reflect_min_meas_value      *float64   `json:"reflect_min_meas_value"`
	Reflect_min_meas_value_ouom *string    `json:"reflect_min_meas_value_ouom"`
	Reflect_min_meas_value_uom  *string    `json:"reflect_min_meas_value_uom"`
	Reflect_std_dev_value       *float64   `json:"reflect_std_dev_value"`
	Refl_random_meas_value      *float64   `json:"refl_random_meas_value"`
	Refl_random_meas_value_ouom *string    `json:"refl_random_meas_value_ouom"`
	Refl_random_meas_value_uom  *string    `json:"refl_random_meas_value_uom"`
	Remark                      *string    `json:"remark"`
	Reworked_percent            *float64   `json:"reworked_percent"`
	Sample_tot_maceral_val      *float64   `json:"sample_tot_maceral_val"`
	Sample_tot_maceral_val_ouom *string    `json:"sample_tot_maceral_val_ouom"`
	Source                      *string    `json:"source"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	Substance_id                *string    `json:"substance_id"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
