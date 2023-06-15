package dto

type Legal_offshore_loc struct {
	Legal_offshore_loc_id  string   `json:"legal_offshore_loc_id" default:"source"`
	Location_type          string   `json:"location_type" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Block_addition         *string  `json:"block_addition" default:""`
	Block_modifier         *string  `json:"block_modifier" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Ew_footage             *float64 `json:"ew_footage" default:""`
	Ew_footage_ouom        *string  `json:"ew_footage_ouom" default:""`
	Ew_start_line          *string  `json:"ew_start_line" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Footage_origin         *string  `json:"footage_origin" default:""`
	Governing_agency_ba_id *string  `json:"governing_agency_ba_id" default:""`
	Grid_block_ew          *string  `json:"grid_block_ew" default:""`
	Grid_block_ns          *string  `json:"grid_block_ns" default:""`
	Grid_block_range       *float64 `json:"grid_block_range" default:""`
	Grid_block_tier        *float64 `json:"grid_block_tier" default:""`
	Ns_footage             *float64 `json:"ns_footage" default:""`
	Ns_footage_ouom        *string  `json:"ns_footage_ouom" default:""`
	Ns_start_line          *string  `json:"ns_start_line" default:""`
	Ocs_num                *string  `json:"ocs_num" default:""`
	Offshore_area_code     *string  `json:"offshore_area_code" default:""`
	Offshore_block         *string  `json:"offshore_block" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Projected_ew_direction *string  `json:"projected_ew_direction" default:""`
	Projected_meridian     *string  `json:"projected_meridian" default:""`
	Projected_ns_direction *string  `json:"projected_ns_direction" default:""`
	Projected_range        *float64 `json:"projected_range" default:""`
	Projected_section      *float64 `json:"projected_section" default:""`
	Projected_township     *float64 `json:"projected_township" default:""`
	Remark                 *string  `json:"remark" default:""`
	Scaled_footage_ind     *string  `json:"scaled_footage_ind" default:""`
	Utm_quadrant           *string  `json:"utm_quadrant" default:""`
	Uwi                    *string  `json:"uwi" default:""`
	Water_bottom_zone      *string  `json:"water_bottom_zone" default:""`
	Well_node_id           *string  `json:"well_node_id" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
