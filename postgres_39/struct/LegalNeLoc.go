package dto

import (
	"time"
)

type Legal_ne_loc struct {
	Legal_ne_loc_id     string     `json:"legal_ne_loc_id"`
	Location_type       string     `json:"location_type"`
	Source              string     `json:"source"`
	Active_ind          *string    `json:"active_ind"`
	Area_id             *string    `json:"area_id"`
	Area_type           *string    `json:"area_type"`
	Coord_system_id     *string    `json:"coord_system_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Ew_footage          *float64   `json:"ew_footage"`
	Ew_footage_ouom     *string    `json:"ew_footage_ouom"`
	Ew_start_line       *string    `json:"ew_start_line"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Footage_origin      *string    `json:"footage_origin"`
	Ne_district         *string    `json:"ne_district"`
	Ne_lot_code         *string    `json:"ne_lot_code"`
	Ne_map_quad_min     *float64   `json:"ne_map_quad_min"`
	Ne_map_quad_name    *string    `json:"ne_map_quad_name"`
	Ne_map_quad_section *string    `json:"ne_map_quad_section"`
	Ne_township_name    *string    `json:"ne_township_name"`
	Ns_footage          *float64   `json:"ns_footage"`
	Ns_footage_ouom     *string    `json:"ns_footage_ouom"`
	Ns_start_line       *string    `json:"ns_start_line"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Reference_latitude  *float64   `json:"reference_latitude"`
	Reference_longitude *float64   `json:"reference_longitude"`
	Remark              *string    `json:"remark"`
	Scaled_footage_ind  *string    `json:"scaled_footage_ind"`
	Uwi                 *string    `json:"uwi"`
	Well_node_id        *string    `json:"well_node_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
