package dto

import (
	"time"
)

type Legal_carter_loc struct {
	Legal_carter_loc_id string     `json:"legal_carter_loc_id"`
	Location_type       string     `json:"location_type"`
	Source              string     `json:"source"`
	Active_ind          *string    `json:"active_ind"`
	Area_id             *string    `json:"area_id"`
	Area_type           *string    `json:"area_type"`
	Carter_range        *float64   `json:"carter_range"`
	Carter_section      *float64   `json:"carter_section"`
	Carter_township     *string    `json:"carter_township"`
	Coord_system_id     *string    `json:"coord_system_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Ew_direction        *string    `json:"ew_direction"`
	Ew_footage          *float64   `json:"ew_footage"`
	Ew_footage_ouom     *string    `json:"ew_footage_ouom"`
	Ew_start_line       *string    `json:"ew_start_line"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Footage_origin      *string    `json:"footage_origin"`
	Map_quad_min        *float64   `json:"map_quad_min"`
	Map_quad_name       *string    `json:"map_quad_name"`
	Ns_direction        *string    `json:"ns_direction"`
	Ns_footage          *float64   `json:"ns_footage"`
	Ns_footage_ouom     *string    `json:"ns_footage_ouom"`
	Ns_start_line       *string    `json:"ns_start_line"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Scaled_footage_ind  *string    `json:"scaled_footage_ind"`
	Spot_code           *string    `json:"spot_code"`
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
