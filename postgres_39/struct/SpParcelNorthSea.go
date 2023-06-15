package dto

import (
	"time"
)

type Sp_parcel_north_sea struct {
	Parcel_north_sea_id    string     `json:"parcel_north_sea_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Block_no               *int       `json:"block_no"`
	Block_suffix           *string    `json:"block_suffix"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Principal_meridian     *string    `json:"principal_meridian"`
	Quadrant               *float64   `json:"quadrant"`
	Quadrant_prefix        *string    `json:"quadrant_prefix"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Subdivision            *string    `json:"subdivision"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
