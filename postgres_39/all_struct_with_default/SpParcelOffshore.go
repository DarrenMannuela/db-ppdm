package dto

type Sp_parcel_offshore struct {
	Parcel_offshore_id     string   `json:"parcel_offshore_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Block_addition         *string  `json:"block_addition" default:""`
	Block_modifier         *string  `json:"block_modifier" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Ew_direction           *string  `json:"ew_direction" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Grid_block_range       *float64 `json:"grid_block_range" default:""`
	Grid_block_tier        *float64 `json:"grid_block_tier" default:""`
	Ns_direction           *string  `json:"ns_direction" default:""`
	Ocs_num                *string  `json:"ocs_num" default:""`
	Offshore_area_code     *string  `json:"offshore_area_code" default:""`
	Offshore_block         *string  `json:"offshore_block" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Reference_plan_num     *string  `json:"reference_plan_num" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Utm_quadrant           *string  `json:"utm_quadrant" default:""`
	Water_bottom_zone      *string  `json:"water_bottom_zone" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
