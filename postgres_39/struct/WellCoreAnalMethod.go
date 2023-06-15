package dto

import (
	"time"
)

type Well_core_anal_method struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Core_id                   string     `json:"core_id"`
	Analysis_obs_no           int        `json:"analysis_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Core_solvent              *string    `json:"core_solvent"`
	Cutting_fluid             *string    `json:"cutting_fluid"`
	Density_analysis          *string    `json:"density_analysis"`
	Drying_method             *string    `json:"drying_method"`
	Drying_temperature        *float64   `json:"drying_temperature"`
	Drying_temperature_ouom   *string    `json:"drying_temperature_ouom"`
	Drying_time_elapsed       *float64   `json:"drying_time_elapsed"`
	Drying_time_elapsed_ouom  *string    `json:"drying_time_elapsed_ouom"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Extract_method            *string    `json:"extract_method"`
	Extract_time_elapsed      *float64   `json:"extract_time_elapsed"`
	Extract_time_elapsed_ouom *string    `json:"extract_time_elapsed_ouom"`
	Fluid_sat_analysis        *string    `json:"fluid_sat_analysis"`
	Permeability_analysis     *string    `json:"permeability_analysis"`
	Porosity_analysis         *string    `json:"porosity_analysis"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
