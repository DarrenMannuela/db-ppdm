package dto

import (
	"time"
)

type Legal_loc_remark struct {
	Legal_loc_remark_id    string     `json:"legal_loc_remark_id"`
	Location_type          string     `json:"location_type"`
	Source                 string     `json:"source"`
	Remark_seq_no          int        `json:"remark_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Land_parcel_type       *string    `json:"land_parcel_type"`
	Legal_carter_loc_id    *string    `json:"legal_carter_loc_id"`
	Legal_congress_loc_id  *string    `json:"legal_congress_loc_id"`
	Legal_dls_loc_id       *string    `json:"legal_dls_loc_id"`
	Legal_fps_loc_id       *string    `json:"legal_fps_loc_id"`
	Legal_geodetic_loc_id  *string    `json:"legal_geodetic_loc_id"`
	Legal_ne_loc_id        *string    `json:"legal_ne_loc_id"`
	Legal_north_sea_loc_id *string    `json:"legal_north_sea_loc_id"`
	Legal_nts_loc_id       *string    `json:"legal_nts_loc_id"`
	Legal_offshore_loc_id  *string    `json:"legal_offshore_loc_id"`
	Legal_ohio_loc_id      *string    `json:"legal_ohio_loc_id"`
	Legal_texas_loc_id     *string    `json:"legal_texas_loc_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Remark_type            *string    `json:"remark_type"`
	Well_node_id           *string    `json:"well_node_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
