package dto

type Sp_mineral_zone struct {
	Spatial_description_id   string  `json:"spatial_description_id" default:"source"`
	Spatial_obs_no           int     `json:"spatial_obs_no" default:"1"`
	Mineral_zone_id          string  `json:"mineral_zone_id" default:"source"`
	Active_ind               *string `json:"active_ind" default:""`
	Base_qualifier           *string `json:"base_qualifier" default:""`
	Base_zone_definition_id  *string `json:"base_zone_definition_id" default:""`
	Deep_right_reversion_ind *string `json:"deep_right_reversion_ind" default:""`
	Effective_date           *string `json:"effective_date" default:""`
	Expiry_date              *string `json:"expiry_date" default:""`
	Inactivation_date        *string `json:"inactivation_date" default:""`
	Metes_and_bounds         *string `json:"metes_and_bounds" default:""`
	Ppdm_guid                *string `json:"ppdm_guid" default:""`
	Remark                   *string `json:"remark" default:""`
	Source                   *string `json:"source" default:""`
	Top_qualifier            *string `json:"top_qualifier" default:""`
	Top_zone_definition_id   *string `json:"top_zone_definition_id" default:""`
	Row_changed_by           *string `json:"row_changed_by" default:""`
	Row_changed_date         *string `json:"row_changed_date" default:""`
	Row_created_by           *string `json:"row_created_by" default:""`
	Row_created_date         *string `json:"row_created_date" default:""`
	Row_effective_date       *string `json:"row_effective_date" default:""`
	Row_expiry_date          *string `json:"row_expiry_date" default:""`
	Row_quality              *string `json:"row_quality" default:""`
}
