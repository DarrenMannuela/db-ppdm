package dto

import (
	"time"
)

type Well_core_sample_anal struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Core_id                    string     `json:"core_id"`
	Analysis_obs_no            int        `json:"analysis_obs_no"`
	Sample_num                 string     `json:"sample_num"`
	Sample_analysis_obs_no     int        `json:"sample_analysis_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Analysis_source            *string    `json:"analysis_source"`
	Bulk_density               *float64   `json:"bulk_density"`
	Bulk_density_ouom          *string    `json:"bulk_density_ouom"`
	Bulk_mass_oil_sat          *float64   `json:"bulk_mass_oil_sat"`
	Bulk_mass_oil_sat_ouom     *string    `json:"bulk_mass_oil_sat_ouom"`
	Bulk_mass_sand_sat         *float64   `json:"bulk_mass_sand_sat"`
	Bulk_mass_sand_sat_ouom    *string    `json:"bulk_mass_sand_sat_ouom"`
	Bulk_mass_water_sat        *float64   `json:"bulk_mass_water_sat"`
	Bulk_mass_water_sat_ouom   *string    `json:"bulk_mass_water_sat_ouom"`
	Bulk_volume_oil_sat        *float64   `json:"bulk_volume_oil_sat"`
	Bulk_volume_water_sat      *float64   `json:"bulk_volume_water_sat"`
	Confine_perm_pressure      *float64   `json:"confine_perm_pressure"`
	Confine_perm_pressure_ouom *string    `json:"confine_perm_pressure_ouom"`
	Confine_por_pressure       *float64   `json:"confine_por_pressure"`
	Confine_por_pressure_ouom  *string    `json:"confine_por_pressure_ouom"`
	Confine_sat_pressure       *float64   `json:"confine_sat_pressure"`
	Confine_sat_pressure_ouom  *string    `json:"confine_sat_pressure_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Effective_porosity         *float64   `json:"effective_porosity"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Gas_sat_volume             *float64   `json:"gas_sat_volume"`
	Grain_density              *float64   `json:"grain_density"`
	Grain_density_ouom         *string    `json:"grain_density_ouom"`
	Grain_mass_oil_sat         *float64   `json:"grain_mass_oil_sat"`
	Grain_mass_oil_sat_ouom    *string    `json:"grain_mass_oil_sat_ouom"`
	Grain_mass_water_sat       *float64   `json:"grain_mass_water_sat"`
	Grain_mass_water_sat_ouom  *string    `json:"grain_mass_water_sat_ouom"`
	Interval_depth             *float64   `json:"interval_depth"`
	Interval_depth_ouom        *string    `json:"interval_depth_ouom"`
	Interval_length            *float64   `json:"interval_length"`
	Interval_length_ouom       *string    `json:"interval_length_ouom"`
	K90                        *float64   `json:"k_90"`
	K90_ouom                   *string    `json:"k_90_ouom"`
	K90_qualifier              *string    `json:"k_90_qualifier"`
	Kmax                       *float64   `json:"kmax"`
	Kmax_ouom                  *string    `json:"kmax_ouom"`
	Kmax_qualifier             *string    `json:"kmax_qualifier"`
	Kvert                      *float64   `json:"kvert"`
	Kvert_ouom                 *string    `json:"kvert_ouom"`
	Kvert_qualifier            *string    `json:"kvert_qualifier"`
	Oil_sat                    *float64   `json:"oil_sat"`
	Pore_volume_gas_sat        *float64   `json:"pore_volume_gas_sat"`
	Pore_volume_oil_sat        *float64   `json:"pore_volume_oil_sat"`
	Pore_volume_water_sat      *float64   `json:"pore_volume_water_sat"`
	Porosity                   *float64   `json:"porosity"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Top_depth                  *float64   `json:"top_depth"`
	Top_depth_ouom             *string    `json:"top_depth_ouom"`
	Water_sat                  *float64   `json:"water_sat"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
