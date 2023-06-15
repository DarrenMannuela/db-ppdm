package dto

import (
	"time"
)

type Well_test_analysis struct {
	Uwi                         string     `json:"uwi"`
	Source                      string     `json:"source"`
	Test_type                   string     `json:"test_type"`
	Run_num                     string     `json:"run_num"`
	Test_num                    string     `json:"test_num"`
	Period_type                 string     `json:"period_type"`
	Period_obs_no               int        `json:"period_obs_no"`
	Analysis_obs_no             int        `json:"analysis_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Bsw                         *float64   `json:"bsw"`
	Completion_obs_no           *int       `json:"completion_obs_no"`
	Condensate_gravity          *float64   `json:"condensate_gravity"`
	Condensate_ratio            *float64   `json:"condensate_ratio"`
	Condensate_ratio_ouom       *string    `json:"condensate_ratio_ouom"`
	Condensate_temperature      *float64   `json:"condensate_temperature"`
	Condensate_temperature_ouom *string    `json:"condensate_temperature_ouom"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Fluid_type                  *string    `json:"fluid_type"`
	Gas_content                 *string    `json:"gas_content"`
	Gas_gravity                 *float64   `json:"gas_gravity"`
	Gor                         *float64   `json:"gor"`
	Gor_ouom                    *string    `json:"gor_ouom"`
	Gwr                         *float64   `json:"gwr"`
	Gwr_ouom                    *string    `json:"gwr_ouom"`
	H2s_percent                 *float64   `json:"h_2_s_percent"`
	Lgr                         *float64   `json:"lgr"`
	Lgr_ouom                    *string    `json:"lgr_ouom"`
	Oil_density                 *float64   `json:"oil_density"`
	Oil_density_ouom            *string    `json:"oil_density_ouom"`
	Oil_gravity                 *float64   `json:"oil_gravity"`
	Oil_temperature             *float64   `json:"oil_temperature"`
	Oil_temperature_ouom        *string    `json:"oil_temperature_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Salinity_type               *string    `json:"salinity_type"`
	Sulphur_percent             *float64   `json:"sulphur_percent"`
	Water_cut                   *float64   `json:"water_cut"`
	Water_resistivity           *float64   `json:"water_resistivity"`
	Water_resistivity_ouom      *string    `json:"water_resistivity_ouom"`
	Water_salinity              *float64   `json:"water_salinity"`
	Water_salinity_ouom         *string    `json:"water_salinity_ouom"`
	Water_temperature           *float64   `json:"water_temperature"`
	Water_temperature_ouom      *string    `json:"water_temperature_ouom"`
	Wor                         *float64   `json:"wor"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
