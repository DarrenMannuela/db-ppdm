package dto

import (
	"time"
)

type Sp_parcel_ne_loc struct {
	Parcel_ne_loc_id       string     `json:"parcel_ne_loc_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ne_district            *string    `json:"ne_district"`
	Ne_lot_code            *string    `json:"ne_lot_code"`
	Ne_map_quad_min        *float64   `json:"ne_map_quad_min"`
	Ne_map_quad_name       *string    `json:"ne_map_quad_name"`
	Ne_map_quad_section    *string    `json:"ne_map_quad_section"`
	Ne_township_name       *string    `json:"ne_township_name"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_plan_num     *string    `json:"reference_plan_num"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
