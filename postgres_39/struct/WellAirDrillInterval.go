package dto

import (
	"time"
)

type Well_air_drill_interval struct {
	Uwi                         string     `json:"uwi"`
	Air_drill_source            string     `json:"air_drill_source"`
	Air_drill_obs_no            int        `json:"air_drill_obs_no"`
	Depth_obs_no                int        `json:"depth_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Air_gas_code                *string    `json:"air_gas_code"`
	Base_depth                  *float64   `json:"base_depth"`
	Base_depth_ouom             *string    `json:"base_depth_ouom"`
	Compressor_rate_volume      *float64   `json:"compressor_rate_volume"`
	Compressor_rate_volume_ouom *string    `json:"compressor_rate_volume_ouom"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Max_interval_pressure       *float64   `json:"max_interval_pressure"`
	Max_interval_pressure_ouom  *string    `json:"max_interval_pressure_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Top_depth                   *float64   `json:"top_depth"`
	Top_depth_ouom              *string    `json:"top_depth_ouom"`
	Water_influx_amount         *float64   `json:"water_influx_amount"`
	Water_influx_amount_ouom    *string    `json:"water_influx_amount_ouom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
