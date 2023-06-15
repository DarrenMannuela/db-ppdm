package dto

type Sp_parcel_carter struct {
	Parcel_carter_id       string   `json:"parcel_carter_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Carter_range           *float64 `json:"carter_range" default:""`
	Carter_section         *float64 `json:"carter_section" default:""`
	Carter_township        *string  `json:"carter_township" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Ew_direction           *string  `json:"ew_direction" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Map_quad_min           *float64 `json:"map_quad_min" default:""`
	Map_quad_name          *string  `json:"map_quad_name" default:""`
	Ns_direction           *string  `json:"ns_direction" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Reference_plan_num     *string  `json:"reference_plan_num" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Spot_code              *string  `json:"spot_code" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
