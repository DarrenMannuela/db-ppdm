package dto

import (
	"time"
)

type Sp_parcel_fps struct {
	Parcel_fps_id          string     `json:"parcel_fps_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Grid                   *string    `json:"grid"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_latitude     *float64   `json:"reference_latitude"`
	Reference_longitude    *float64   `json:"reference_longitude"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Section                *int       `json:"section"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Unit                   *string    `json:"unit"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
