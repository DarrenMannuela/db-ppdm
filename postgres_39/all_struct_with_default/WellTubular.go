package dto

type Well_tubular struct {
	Uwi                        string   `json:"uwi" default:"source"`
	Source                     string   `json:"source" default:"source"`
	Tubing_type                string   `json:"tubing_type" default:"source"`
	Tubing_obs_no              int      `json:"tubing_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Base_depth                 *float64 `json:"base_depth" default:""`
	Base_depth_ouom            *string  `json:"base_depth_ouom" default:""`
	Base_strat_unit_id         *string  `json:"base_strat_unit_id" default:""`
	Cat_equip_id               *string  `json:"cat_equip_id" default:""`
	Collar_type                *string  `json:"collar_type" default:""`
	Coupling_type              *string  `json:"coupling_type" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Equipment_id               *string  `json:"equipment_id" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Hole_size                  *float64 `json:"hole_size" default:""`
	Hole_size_ouom             *string  `json:"hole_size_ouom" default:""`
	Hung_top_depth             *float64 `json:"hung_top_depth" default:""`
	Hung_top_depth_ouom        *string  `json:"hung_top_depth_ouom" default:""`
	Inside_diameter            *float64 `json:"inside_diameter" default:""`
	Inside_diameter_ouom       *string  `json:"inside_diameter_ouom" default:""`
	Joint_count                *int     `json:"joint_count" default:""`
	Left_in_hole_length        *float64 `json:"left_in_hole_length" default:""`
	Left_in_hole_length_ouom   *string  `json:"left_in_hole_length_ouom" default:""`
	Liner_type                 *string  `json:"liner_type" default:""`
	Manufacturer_ba_id         *string  `json:"manufacturer_ba_id" default:""`
	Mixed_string_ind           *string  `json:"mixed_string_ind" default:""`
	Observation_date           *string  `json:"observation_date" default:""`
	Outside_diameter           *float64 `json:"outside_diameter" default:""`
	Outside_diameter_desc      *string  `json:"outside_diameter_desc" default:""`
	Outside_diameter_ouom      *string  `json:"outside_diameter_ouom" default:""`
	Packer_set_depth           *float64 `json:"packer_set_depth" default:""`
	Packer_set_depth_ouom      *string  `json:"packer_set_depth_ouom" default:""`
	Period_obs_no              *int     `json:"period_obs_no" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Pulled_length              *float64 `json:"pulled_length" default:""`
	Pulled_length_ouom         *string  `json:"pulled_length_ouom" default:""`
	Remark                     *string  `json:"remark" default:""`
	Reported_base_tvd          *float64 `json:"reported_base_tvd" default:""`
	Reported_base_tvd_ouom     *string  `json:"reported_base_tvd_ouom" default:""`
	Reported_top_tvd           *float64 `json:"reported_top_tvd" default:""`
	Reported_top_tvd_ouom      *string  `json:"reported_top_tvd_ouom" default:""`
	Sea_floor_penetration      *float64 `json:"sea_floor_penetration" default:""`
	Sea_floor_penetration_ouom *string  `json:"sea_floor_penetration_ouom" default:""`
	Shoe_depth                 *float64 `json:"shoe_depth" default:""`
	Shoe_depth_ouom            *string  `json:"shoe_depth_ouom" default:""`
	Steel_spec                 *string  `json:"steel_spec" default:""`
	Strat_name_set_id          *string  `json:"strat_name_set_id" default:""`
	Top_strat_unit_id          *string  `json:"top_strat_unit_id" default:""`
	Torque                     *float64 `json:"torque" default:""`
	Torque_ouom                *string  `json:"torque_ouom" default:""`
	Tubing_density             *float64 `json:"tubing_density" default:""`
	Tubing_density_ouom        *string  `json:"tubing_density_ouom" default:""`
	Tubing_grade               *string  `json:"tubing_grade" default:""`
	Tubing_strength            *float64 `json:"tubing_strength" default:""`
	Tubing_strength_ouom       *string  `json:"tubing_strength_ouom" default:""`
	Tubing_weight              *float64 `json:"tubing_weight" default:""`
	Tubing_weight_ouom         *string  `json:"tubing_weight_ouom" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
