package dto

import (
	"time"
)

type Sf_landing struct {
	Sf_subtype              string     `json:"sf_subtype"`
	Support_facility_id     string     `json:"support_facility_id"`
	Activation_freq_desc    *string    `json:"activation_freq_desc"`
	Activation_tone_desc    *string    `json:"activation_tone_desc"`
	Active_ind              *string    `json:"active_ind"`
	Airpspace_desc          *string    `json:"airpspace_desc"`
	Altitude                *float64   `json:"altitude"`
	Altitude_ouom           *string    `json:"altitude_ouom"`
	Approach_direction      *string    `json:"approach_direction"`
	Area_id                 *string    `json:"area_id"`
	Area_size               *float64   `json:"area_size"`
	Area_size_ouom          *string    `json:"area_size_ouom"`
	Area_type               *string    `json:"area_type"`
	Communication_freq      *float64   `json:"communication_freq"`
	Communication_freq_desc *string    `json:"communication_freq_desc"`
	Depart_direction        *string    `json:"depart_direction"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Fuel_avail_desc         *string    `json:"fuel_avail_desc"`
	Landing_facility_code   *string    `json:"landing_facility_code"`
	Landing_type            *string    `json:"landing_type"`
	Lighting_avail_ind      *string    `json:"lighting_avail_ind"`
	Lighting_cycle_desc     *string    `json:"lighting_cycle_desc"`
	Lighting_desc           *string    `json:"lighting_desc"`
	Max_allow_mass          *float64   `json:"max_allow_mass"`
	Max_allow_mass_desc     *string    `json:"max_allow_mass_desc"`
	Max_allow_mass_ouom     *string    `json:"max_allow_mass_ouom"`
	Perm_obstacle_desc      *string    `json:"perm_obstacle_desc"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Radio_call_name         *string    `json:"radio_call_name"`
	Radio_channel           *string    `json:"radio_channel"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Special_procedure_desc  *string    `json:"special_procedure_desc"`
	Surface_desc            *string    `json:"surface_desc"`
	Surface_type            *string    `json:"surface_type"`
	Weather_info_desc       *string    `json:"weather_info_desc"`
	Windcone_ind            *string    `json:"windcone_ind"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
