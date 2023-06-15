package dto

import (
	"time"
)

type Well_tubular struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Tubing_type                string     `json:"tubing_type"`
	Tubing_obs_no              int        `json:"tubing_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Base_depth                 *float64   `json:"base_depth"`
	Base_depth_ouom            *string    `json:"base_depth_ouom"`
	Base_strat_unit_id         *string    `json:"base_strat_unit_id"`
	Cat_equip_id               *string    `json:"cat_equip_id"`
	Collar_type                *string    `json:"collar_type"`
	Coupling_type              *string    `json:"coupling_type"`
	Effective_date             *time.Time `json:"effective_date"`
	Equipment_id               *string    `json:"equipment_id"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Hole_size                  *float64   `json:"hole_size"`
	Hole_size_ouom             *string    `json:"hole_size_ouom"`
	Hung_top_depth             *float64   `json:"hung_top_depth"`
	Hung_top_depth_ouom        *string    `json:"hung_top_depth_ouom"`
	Inside_diameter            *float64   `json:"inside_diameter"`
	Inside_diameter_ouom       *string    `json:"inside_diameter_ouom"`
	Joint_count                *int       `json:"joint_count"`
	Left_in_hole_length        *float64   `json:"left_in_hole_length"`
	Left_in_hole_length_ouom   *string    `json:"left_in_hole_length_ouom"`
	Liner_type                 *string    `json:"liner_type"`
	Manufacturer_ba_id         *string    `json:"manufacturer_ba_id"`
	Mixed_string_ind           *string    `json:"mixed_string_ind"`
	Observation_date           *time.Time `json:"observation_date"`
	Outside_diameter           *float64   `json:"outside_diameter"`
	Outside_diameter_desc      *string    `json:"outside_diameter_desc"`
	Outside_diameter_ouom      *string    `json:"outside_diameter_ouom"`
	Packer_set_depth           *float64   `json:"packer_set_depth"`
	Packer_set_depth_ouom      *string    `json:"packer_set_depth_ouom"`
	Period_obs_no              *int       `json:"period_obs_no"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Pulled_length              *float64   `json:"pulled_length"`
	Pulled_length_ouom         *string    `json:"pulled_length_ouom"`
	Remark                     *string    `json:"remark"`
	Reported_base_tvd          *float64   `json:"reported_base_tvd"`
	Reported_base_tvd_ouom     *string    `json:"reported_base_tvd_ouom"`
	Reported_top_tvd           *float64   `json:"reported_top_tvd"`
	Reported_top_tvd_ouom      *string    `json:"reported_top_tvd_ouom"`
	Sea_floor_penetration      *float64   `json:"sea_floor_penetration"`
	Sea_floor_penetration_ouom *string    `json:"sea_floor_penetration_ouom"`
	Shoe_depth                 *float64   `json:"shoe_depth"`
	Shoe_depth_ouom            *string    `json:"shoe_depth_ouom"`
	Steel_spec                 *string    `json:"steel_spec"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Top_strat_unit_id          *string    `json:"top_strat_unit_id"`
	Torque                     *float64   `json:"torque"`
	Torque_ouom                *string    `json:"torque_ouom"`
	Tubing_density             *float64   `json:"tubing_density"`
	Tubing_density_ouom        *string    `json:"tubing_density_ouom"`
	Tubing_grade               *string    `json:"tubing_grade"`
	Tubing_strength            *float64   `json:"tubing_strength"`
	Tubing_strength_ouom       *string    `json:"tubing_strength_ouom"`
	Tubing_weight              *float64   `json:"tubing_weight"`
	Tubing_weight_ouom         *string    `json:"tubing_weight_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
