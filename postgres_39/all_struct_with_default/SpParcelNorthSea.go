package dto

type Sp_parcel_north_sea struct {
	Parcel_north_sea_id    string   `json:"parcel_north_sea_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Block_no               *int     `json:"block_no" default:""`
	Block_suffix           *string  `json:"block_suffix" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Principal_meridian     *string  `json:"principal_meridian" default:""`
	Quadrant               *float64 `json:"quadrant" default:""`
	Quadrant_prefix        *string  `json:"quadrant_prefix" default:""`
	Reference_plan_num     *string  `json:"reference_plan_num" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Subdivision            *string  `json:"subdivision" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
