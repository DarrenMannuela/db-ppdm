package dto

import (
	"time"
)

type Anl_qgf_tsf struct {
	Analysis_id              string     `json:"analysis_id"`
	Anl_source               string     `json:"anl_source"`
	Measurement_obs_no       int        `json:"measurement_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Emission_intensity       *float64   `json:"emission_intensity"`
	Emission_intensity_uom   *string    `json:"emission_intensity_uom"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Preferred_ind            *string    `json:"preferred_ind"`
	Problem_ind              *string    `json:"problem_ind"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Step_seq_no              *int       `json:"step_seq_no"`
	Wavelength_emission      *float64   `json:"wavelength_emission"`
	Wavelength_emission_uom  *string    `json:"wavelength_emission_uom"`
	Wavelenth_excitation     *float64   `json:"wavelenth_excitation"`
	Wavelenth_excitation_uom *string    `json:"wavelenth_excitation_uom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
