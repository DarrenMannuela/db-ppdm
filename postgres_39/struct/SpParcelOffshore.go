package dto

import (
	"time"
)

type Sp_parcel_offshore struct {
	Parcel_offshore_id     string     `json:"parcel_offshore_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Block_addition         *string    `json:"block_addition"`
	Block_modifier         *string    `json:"block_modifier"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Ew_direction           *string    `json:"ew_direction"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Grid_block_range       *float64   `json:"grid_block_range"`
	Grid_block_tier        *float64   `json:"grid_block_tier"`
	Ns_direction           *string    `json:"ns_direction"`
	Ocs_num                *string    `json:"ocs_num"`
	Offshore_area_code     *string    `json:"offshore_area_code"`
	Offshore_block         *string    `json:"offshore_block"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Utm_quadrant           *string    `json:"utm_quadrant"`
	Water_bottom_zone      *string    `json:"water_bottom_zone"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
