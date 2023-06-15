package dto

import (
	"time"
)

type Strat_field_section struct {
	Field_station_id        string     `json:"field_station_id"`
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
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Fault_heave             *float64   `json:"fault_heave"`
	Fault_throw             *float64   `json:"fault_throw"`
	Interpretation_version  *string    `json:"interpretation_version"`
	Interpreter             *string    `json:"interpreter"`
	Lithology               *string    `json:"lithology"`
	Missing_section         *float64   `json:"missing_section"`
	Missing_strat_type      *string    `json:"missing_strat_type"`
	Ordinal_seq_no          *int       `json:"ordinal_seq_no"`
	Overturned_ind          *string    `json:"overturned_ind"`
	Pick_date               *time.Time `json:"pick_date"`
	Pick_location           *string    `json:"pick_location"`
	Pick_qualifier          *string    `json:"pick_qualifier"`
	Pick_qualif_reason      *string    `json:"pick_qualif_reason"`
	Pick_quality            *string    `json:"pick_quality"`
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
	Strat_unit_md           *float64   `json:"strat_unit_md"`
	Strat_unit_md_ouom      *string    `json:"strat_unit_md_ouom"`
	Strat_unit_tvd          *float64   `json:"strat_unit_tvd"`
	Strike                  *float64   `json:"strike"`
	Sw_application_id       *string    `json:"sw_application_id"`
	True_thickness_ind      *string    `json:"true_thickness_ind"`
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
