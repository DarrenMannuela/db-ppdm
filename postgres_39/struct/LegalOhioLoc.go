package dto

import (
	"time"
)

type Legal_ohio_loc struct {
	Legal_ohio_loc_id          string     `json:"legal_ohio_loc_id"`
	Location_type              string     `json:"location_type"`
	Source                     string     `json:"source"`
	Active_ind                 *string    `json:"active_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_type                  *string    `json:"area_type"`
	Coord_system_id            *string    `json:"coord_system_id"`
	Effective_date             *time.Time `json:"effective_date"`
	Ew_footage                 *float64   `json:"ew_footage"`
	Ew_footage_ouom            *string    `json:"ew_footage_ouom"`
	Ew_start_line              *string    `json:"ew_start_line"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Footage_origin             *string    `json:"footage_origin"`
	Map_quad_min               *float64   `json:"map_quad_min"`
	Map_quad_name              *string    `json:"map_quad_name"`
	Ns_footage                 *float64   `json:"ns_footage"`
	Ns_footage_ouom            *string    `json:"ns_footage_ouom"`
	Ns_start_line              *string    `json:"ns_start_line"`
	Ohio_allotment             *string    `json:"ohio_allotment"`
	Ohio_division              *string    `json:"ohio_division"`
	Ohio_fraction              *string    `json:"ohio_fraction"`
	Ohio_land_subdivision_name *string    `json:"ohio_land_subdivision_name"`
	Ohio_other_subdivision     *string    `json:"ohio_other_subdivision"`
	Ohio_quarter_township      *string    `json:"ohio_quarter_township"`
	Ohio_range                 *float64   `json:"ohio_range"`
	Ohio_range_dir             *string    `json:"ohio_range_dir"`
	Ohio_township              *float64   `json:"ohio_township"`
	Ohio_township_dir          *string    `json:"ohio_township_dir"`
	Ohio_township_name         *string    `json:"ohio_township_name"`
	Ohio_tract                 *string    `json:"ohio_tract"`
	Ohio_twp_lot_code          *string    `json:"ohio_twp_lot_code"`
	Ohio_twp_section_code      *string    `json:"ohio_twp_section_code"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Principal_meridian         *string    `json:"principal_meridian"`
	Remark                     *string    `json:"remark"`
	Scaled_footage_ind         *string    `json:"scaled_footage_ind"`
	Spot_code                  *string    `json:"spot_code"`
	Uwi                        *string    `json:"uwi"`
	Well_node_id               *string    `json:"well_node_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
