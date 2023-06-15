package dto

type Well_air_drill_interval struct {
	Uwi                         string   `json:"uwi" default:"source"`
	Air_drill_source            string   `json:"air_drill_source" default:"source"`
	Air_drill_obs_no            int      `json:"air_drill_obs_no" default:"1"`
	Depth_obs_no                int      `json:"depth_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Air_gas_code                *string  `json:"air_gas_code" default:""`
	Base_depth                  *float64 `json:"base_depth" default:""`
	Base_depth_ouom             *string  `json:"base_depth_ouom" default:""`
	Compressor_rate_volume      *float64 `json:"compressor_rate_volume" default:""`
	Compressor_rate_volume_ouom *string  `json:"compressor_rate_volume_ouom" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Max_interval_pressure       *float64 `json:"max_interval_pressure" default:""`
	Max_interval_pressure_ouom  *string  `json:"max_interval_pressure_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Top_depth                   *float64 `json:"top_depth" default:""`
	Top_depth_ouom              *string  `json:"top_depth_ouom" default:""`
	Water_influx_amount         *float64 `json:"water_influx_amount" default:""`
	Water_influx_amount_ouom    *string  `json:"water_influx_amount_ouom" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
