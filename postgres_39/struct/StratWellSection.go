package dto

import (
	"time"
)

type Strat_well_section struct {
	Uwi                     string     `json:"uwi"`
	Strat_name_set_id       string     `json:"strat_name_set_id"`
	Strat_unit_id           string     `json:"strat_unit_id"`
	Interp_id               string     `json:"interp_id"`
	Active_ind              *string    `json:"active_ind"`
	Area_id                 *string    `json:"area_id"`
	Area_type               *string    `json:"area_type"`
	Azimuth_north_type      *string    `json:"azimuth_north_type"`
	Certified_ind           *string    `json:"certified_ind"`
	Conformity_relationship *string    `json:"conformity_relationship"`
	Dip_angle               *float64   `json:"dip_angle"`
	Dip_direction           *string    `json:"dip_direction"`
	Dominant_lithology      *string    `json:"dominant_lithology"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Fault_heave             *float64   `json:"fault_heave"`
	Fault_throw             *float64   `json:"fault_throw"`
	Interpreter             *string    `json:"interpreter"`
	Missing_section         *float64   `json:"missing_section"`
	Missing_strat_type      *string    `json:"missing_strat_type"`
	Ordinal_seq_no          *int       `json:"ordinal_seq_no"`
	Overturned_ind          *string    `json:"overturned_ind"`
	Pick_date               *time.Time `json:"pick_date"`
	Pick_depth              *float64   `json:"pick_depth"`
	Pick_depth_ouom         *string    `json:"pick_depth_ouom"`
	Pick_location           *string    `json:"pick_location"`
	Pick_qualifier          *string    `json:"pick_qualifier"`
	Pick_qualif_reason      *string    `json:"pick_qualif_reason"`
	Pick_quality            *string    `json:"pick_quality"`
	Pick_tvd                *float64   `json:"pick_tvd"`
	Pick_version_type       *string    `json:"pick_version_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Preferred_pick_ind      *string    `json:"preferred_pick_ind"`
	Remark                  *string    `json:"remark"`
	Repeat_section          *float64   `json:"repeat_section"`
	Repeat_strat_occur_no   *int       `json:"repeat_strat_occur_no"`
	Repeat_strat_type       *string    `json:"repeat_strat_type"`
	Source                  *string    `json:"source"`
	Source_document_id      *string    `json:"source_document_id"`
	Strat_interpret_method  *string    `json:"strat_interpret_method"`
	Strike                  *float64   `json:"strike"`
	Sw_application_id       *string    `json:"sw_application_id"`
	Tvd_method              *string    `json:"tvd_method"`
	Version_obs_no          *int       `json:"version_obs_no"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
