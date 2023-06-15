package dto

type Sp_zone_substance struct {
	Spatial_description_id string   `json:"spatial_description_id" default:"source"`
	Spatial_obs_no         int      `json:"spatial_obs_no" default:"1"`
	Mineral_zone_id        string   `json:"mineral_zone_id" default:"source"`
	Substance_id           string   `json:"substance_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Base_depth             *float64 `json:"base_depth" default:""`
	Base_depth_ouom        *string  `json:"base_depth_ouom" default:""`
	Base_qualifier         *string  `json:"base_qualifier" default:""`
	Base_strat_unit_id     *string  `json:"base_strat_unit_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Excluded_ind           *string  `json:"excluded_ind" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Included_ind           *string  `json:"included_ind" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Strat_name_set_id      *string  `json:"strat_name_set_id" default:""`
	Top_depth              *float64 `json:"top_depth" default:""`
	Top_depth_ouom         *string  `json:"top_depth_ouom" default:""`
	Top_qualifier          *string  `json:"top_qualifier" default:""`
	Top_strat_unit_id      *string  `json:"top_strat_unit_id" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
