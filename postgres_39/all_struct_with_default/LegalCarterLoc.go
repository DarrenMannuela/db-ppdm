package dto

type Legal_carter_loc struct {
	Legal_carter_loc_id string   `json:"legal_carter_loc_id" default:"source"`
	Location_type       string   `json:"location_type" default:"source"`
	Source              string   `json:"source" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Area_id             *string  `json:"area_id" default:""`
	Area_type           *string  `json:"area_type" default:""`
	Carter_range        *float64 `json:"carter_range" default:""`
	Carter_section      *float64 `json:"carter_section" default:""`
	Carter_township     *string  `json:"carter_township" default:""`
	Coord_system_id     *string  `json:"coord_system_id" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Ew_direction        *string  `json:"ew_direction" default:""`
	Ew_footage          *float64 `json:"ew_footage" default:""`
	Ew_footage_ouom     *string  `json:"ew_footage_ouom" default:""`
	Ew_start_line       *string  `json:"ew_start_line" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Footage_origin      *string  `json:"footage_origin" default:""`
	Map_quad_min        *float64 `json:"map_quad_min" default:""`
	Map_quad_name       *string  `json:"map_quad_name" default:""`
	Ns_direction        *string  `json:"ns_direction" default:""`
	Ns_footage          *float64 `json:"ns_footage" default:""`
	Ns_footage_ouom     *string  `json:"ns_footage_ouom" default:""`
	Ns_start_line       *string  `json:"ns_start_line" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Scaled_footage_ind  *string  `json:"scaled_footage_ind" default:""`
	Spot_code           *string  `json:"spot_code" default:""`
	Uwi                 *string  `json:"uwi" default:""`
	Well_node_id        *string  `json:"well_node_id" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
