package dto

import (
	"time"
)

type Well_test_flow struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Test_type                 string     `json:"test_type"`
	Run_num                   string     `json:"run_num"`
	Test_num                  string     `json:"test_num"`
	Period_type               string     `json:"period_type"`
	Period_obs_no             int        `json:"period_obs_no"`
	Fluid_type                string     `json:"fluid_type"`
	Active_ind                *string    `json:"active_ind"`
	Basic_sediment_ind        *string    `json:"basic_sediment_ind"`
	Blow_desc                 *string    `json:"blow_desc"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Gas_riser_diameter        *float64   `json:"gas_riser_diameter"`
	Gas_riser_diameter_ouom   *string    `json:"gas_riser_diameter_ouom"`
	Max_fluid_rate            *float64   `json:"max_fluid_rate"`
	Max_fluid_rate_ouom       *string    `json:"max_fluid_rate_ouom"`
	Max_fluid_rate_uom        *string    `json:"max_fluid_rate_uom"`
	Max_surface_pressure      *float64   `json:"max_surface_pressure"`
	Max_surface_pressure_ouom *string    `json:"max_surface_pressure_ouom"`
	Measurement_technique     *string    `json:"measurement_technique"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Shakeout_percent          *float64   `json:"shakeout_percent"`
	Tts_time_elapsed          *float64   `json:"tts_time_elapsed"`
	Tts_time_elapsed_ouom     *string    `json:"tts_time_elapsed_ouom"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
