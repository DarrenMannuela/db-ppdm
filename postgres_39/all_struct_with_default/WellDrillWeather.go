package dto

type Well_drill_weather struct {
	Uwi                string   `json:"uwi" default:"source"`
	Period_obs_no      int      `json:"period_obs_no" default:"1"`
	Weather_obs_no     int      `json:"weather_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Heave              *float64 `json:"heave" default:""`
	Pitch              *float64 `json:"pitch" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Recorded_time      *string  `json:"recorded_time" default:""`
	Recorded_timezone  *string  `json:"recorded_timezone" default:""`
	Remark             *string  `json:"remark" default:""`
	Road_condition     *string  `json:"road_condition" default:""`
	Roll               *float64 `json:"roll" default:""`
	Source             *string  `json:"source" default:""`
	Swell_height       *float64 `json:"swell_height" default:""`
	Temperature        *float64 `json:"temperature" default:""`
	Temperature_ouom   *string  `json:"temperature_ouom" default:""`
	Water_condition    *string  `json:"water_condition" default:""`
	Wave_height        *float64 `json:"wave_height" default:""`
	Weather_condition  *string  `json:"weather_condition" default:""`
	Wind_direction     *string  `json:"wind_direction" default:""`
	Wind_strength      *string  `json:"wind_strength" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
