package dto

import (
	"time"
)

type Legal_offshore_loc struct {
	Legal_offshore_loc_id  string     `json:"legal_offshore_loc_id"`
	Location_type          string     `json:"location_type"`
	Source                 string     `json:"source"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Block_addition         *string    `json:"block_addition"`
	Block_modifier         *string    `json:"block_modifier"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Ew_footage             *float64   `json:"ew_footage"`
	Ew_footage_ouom        *string    `json:"ew_footage_ouom"`
	Ew_start_line          *string    `json:"ew_start_line"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Footage_origin         *string    `json:"footage_origin"`
	Governing_agency_ba_id *string    `json:"governing_agency_ba_id"`
	Grid_block_ew          *string    `json:"grid_block_ew"`
	Grid_block_ns          *string    `json:"grid_block_ns"`
	Grid_block_range       *float64   `json:"grid_block_range"`
	Grid_block_tier        *float64   `json:"grid_block_tier"`
	Ns_footage             *float64   `json:"ns_footage"`
	Ns_footage_ouom        *string    `json:"ns_footage_ouom"`
	Ns_start_line          *string    `json:"ns_start_line"`
	Ocs_num                *string    `json:"ocs_num"`
	Offshore_area_code     *string    `json:"offshore_area_code"`
	Offshore_block         *string    `json:"offshore_block"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Projected_ew_direction *string    `json:"projected_ew_direction"`
	Projected_meridian     *string    `json:"projected_meridian"`
	Projected_ns_direction *string    `json:"projected_ns_direction"`
	Projected_range        *float64   `json:"projected_range"`
	Projected_section      *float64   `json:"projected_section"`
	Projected_township     *float64   `json:"projected_township"`
	Remark                 *string    `json:"remark"`
	Scaled_footage_ind     *string    `json:"scaled_footage_ind"`
	Utm_quadrant           *string    `json:"utm_quadrant"`
	Uwi                    *string    `json:"uwi"`
	Water_bottom_zone      *string    `json:"water_bottom_zone"`
	Well_node_id           *string    `json:"well_node_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
