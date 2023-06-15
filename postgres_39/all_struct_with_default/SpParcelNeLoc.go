package dto

type Sp_parcel_ne_loc struct {
	Parcel_ne_loc_id       string   `json:"parcel_ne_loc_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Ne_district            *string  `json:"ne_district" default:""`
	Ne_lot_code            *string  `json:"ne_lot_code" default:""`
	Ne_map_quad_min        *float64 `json:"ne_map_quad_min" default:""`
	Ne_map_quad_name       *string  `json:"ne_map_quad_name" default:""`
	Ne_map_quad_section    *string  `json:"ne_map_quad_section" default:""`
	Ne_township_name       *string  `json:"ne_township_name" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Reference_plan_num     *string  `json:"reference_plan_num" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Spatial_description_id *string  `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int     `json:"spatial_obs_no" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
