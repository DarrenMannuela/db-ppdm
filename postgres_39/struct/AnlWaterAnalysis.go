package dto

import (
	"time"
)

type Anl_water_analysis struct {
	Analysis_id                string     `json:"analysis_id"`
	Anl_source                 string     `json:"anl_source"`
	Water_analysis_obs_no      int        `json:"water_analysis_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Barium                     *float64   `json:"barium"`
	Barium_ouom                *string    `json:"barium_ouom"`
	Bicarbonate                *float64   `json:"bicarbonate"`
	Bicarbonate_ouom           *string    `json:"bicarbonate_ouom"`
	Boron                      *float64   `json:"boron"`
	Boron_ouom                 *string    `json:"boron_ouom"`
	Bromine                    *float64   `json:"bromine"`
	Bromine_ouom               *string    `json:"bromine_ouom"`
	Calcium                    *float64   `json:"calcium"`
	Calcium_ouom               *string    `json:"calcium_ouom"`
	Carbonate                  *float64   `json:"carbonate"`
	Carbonate_ouom             *string    `json:"carbonate_ouom"`
	Chlorine                   *float64   `json:"chlorine"`
	Chlorine_ouom              *string    `json:"chlorine_ouom"`
	Conductivity               *float64   `json:"conductivity"`
	Conductivity_ouom          *string    `json:"conductivity_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Hydrogen_sulphide          *float64   `json:"hydrogen_sulphide"`
	Hydrogen_sulphide_ouom     *string    `json:"hydrogen_sulphide_ouom"`
	Hydroxide                  *float64   `json:"hydroxide"`
	Hydroxide_ouom             *string    `json:"hydroxide_ouom"`
	Iodine                     *float64   `json:"iodine"`
	Iodine_ouom                *string    `json:"iodine_ouom"`
	Ion_ouom                   *string    `json:"ion_ouom"`
	Ion_uom                    *string    `json:"ion_uom"`
	Iron                       *float64   `json:"iron"`
	Iron_ouom                  *string    `json:"iron_ouom"`
	Magnesium                  *float64   `json:"magnesium"`
	Magnesium_ouom             *string    `json:"magnesium_ouom"`
	Organics_desc              *string    `json:"organics_desc"`
	Ph_observed                *float64   `json:"ph_observed"`
	Ph_temp                    *float64   `json:"ph_temp"`
	Ph_temp_ouom               *string    `json:"ph_temp_ouom"`
	Potassium                  *float64   `json:"potassium"`
	Potassium_ouom             *string    `json:"potassium_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_ind              *string    `json:"preferred_ind"`
	Problem_ind                *string    `json:"problem_ind"`
	Recovery_desc              *string    `json:"recovery_desc"`
	Refractive_index           *float64   `json:"refractive_index"`
	Refractive_index_temp      *float64   `json:"refractive_index_temp"`
	Refractive_index_temp_ouom *string    `json:"refractive_index_temp_ouom"`
	Remark                     *string    `json:"remark"`
	Reservoir_temperature      *float64   `json:"reservoir_temperature"`
	Reservoir_temperature_ouom *string    `json:"reservoir_temperature_ouom"`
	Rw_temperature             *float64   `json:"rw_temperature"`
	Rw_temperature_ouom        *string    `json:"rw_temperature_ouom"`
	Sample_water_type          *string    `json:"sample_water_type"`
	Sodium                     *float64   `json:"sodium"`
	Sodium_ouom                *string    `json:"sodium_ouom"`
	Sodium_potassium           *float64   `json:"sodium_potassium"`
	Sodium_potassium_ouom      *string    `json:"sodium_potassium_ouom"`
	Source                     *string    `json:"source"`
	Step_seq_no                *int       `json:"step_seq_no"`
	Strontium                  *float64   `json:"strontium"`
	Strontium_ouom             *string    `json:"strontium_ouom"`
	Sulphate                   *float64   `json:"sulphate"`
	Sulphate_ouom              *string    `json:"sulphate_ouom"`
	Water_density              *float64   `json:"water_density"`
	Water_density_ouom         *string    `json:"water_density_ouom"`
	Water_ph                   *float64   `json:"water_ph"`
	Water_resistivity          *float64   `json:"water_resistivity"`
	Water_resistivity_ouom     *string    `json:"water_resistivity_ouom"`
	Water_type                 *string    `json:"water_type"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
