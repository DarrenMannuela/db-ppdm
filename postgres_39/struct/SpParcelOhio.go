package dto

import (
	"time"
)

type Sp_parcel_ohio struct {
	Parcel_ohio_id             string     `json:"parcel_ohio_id"`
	Active_ind                 *string    `json:"active_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_type                  *string    `json:"area_type"`
	Coord_system_id            *string    `json:"coord_system_id"`
	Description                *string    `json:"description"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Jurisdiction               *string    `json:"jurisdiction"`
	Map_quad_min               *float64   `json:"map_quad_min"`
	Map_quad_name              *string    `json:"map_quad_name"`
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
	Reference_plan_num         *string    `json:"reference_plan_num"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Spatial_description_id     *string    `json:"spatial_description_id"`
	Spatial_obs_no             *int       `json:"spatial_obs_no"`
	Spot_code                  *string    `json:"spot_code"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
