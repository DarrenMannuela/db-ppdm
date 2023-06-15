package dto

import (
	"time"
)

type Hse_incident_weather struct {
	Incident_id        string     `json:"incident_id"`
	Weather_obs_no     int        `json:"weather_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Heave              *float64   `json:"heave"`
	Pitch              *float64   `json:"pitch"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Recorded_time      *time.Time `json:"recorded_time"`
	Recorded_timezone  *string    `json:"recorded_timezone"`
	Remark             *string    `json:"remark"`
	Road_condition     *string    `json:"road_condition"`
	Roll               *float64   `json:"roll"`
	Source             *string    `json:"source"`
	Swell_height       *float64   `json:"swell_height"`
	Temperature        *float64   `json:"temperature"`
	Temperature_ouom   *string    `json:"temperature_ouom"`
	Water_condition    *string    `json:"water_condition"`
	Wave_height        *float64   `json:"wave_height"`
	Weather_condition  *string    `json:"weather_condition"`
	Wind_direction     *string    `json:"wind_direction"`
	Wind_strength      *string    `json:"wind_strength"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
