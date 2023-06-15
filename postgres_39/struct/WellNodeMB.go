package dto

import (
	"time"
)

type Well_node_m_b struct {
	Node_id               string     `json:"node_id"`
	Source                string     `json:"source"`
	Active_ind            *string    `json:"active_ind"`
	Dls_road_allowance_id *string    `json:"dls_road_allowance_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Ew_direction          *string    `json:"ew_direction"`
	Ew_distance           *float64   `json:"ew_distance"`
	Ew_distance_ouom      *string    `json:"ew_distance_ouom"`
	Ew_start_line         *string    `json:"ew_start_line"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Location_type         *string    `json:"location_type"`
	Ns_direction          *string    `json:"ns_direction"`
	Ns_distance           *float64   `json:"ns_distance"`
	Ns_distance_ouom      *string    `json:"ns_distance_ouom"`
	Ns_start_line         *string    `json:"ns_start_line"`
	Orientation           *string    `json:"orientation"`
	Parcel_carter_id      *string    `json:"parcel_carter_id"`
	Parcel_congress_id    *string    `json:"parcel_congress_id"`
	Parcel_dls_id         *string    `json:"parcel_dls_id"`
	Parcel_fps_id         *string    `json:"parcel_fps_id"`
	Parcel_libya_id       *string    `json:"parcel_libya_id"`
	Parcel_ne_loc_id      *string    `json:"parcel_ne_loc_id"`
	Parcel_north_sea_id   *string    `json:"parcel_north_sea_id"`
	Parcel_nts_id         *string    `json:"parcel_nts_id"`
	Parcel_offshore_id    *string    `json:"parcel_offshore_id"`
	Parcel_ohio_id        *string    `json:"parcel_ohio_id"`
	Parcel_pbl_id         *string    `json:"parcel_pbl_id"`
	Parcel_texas          *string    `json:"parcel_texas"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Reference_loc         *string    `json:"reference_loc"`
	Remark                *string    `json:"remark"`
	Surface_loc           *string    `json:"surface_loc"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
