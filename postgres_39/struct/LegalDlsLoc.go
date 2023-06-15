package dto

import (
	"time"
)

type Legal_dls_loc struct {
	Legal_dls_loc_id       string     `json:"legal_dls_loc_id"`
	Location_type          string     `json:"location_type"`
	Source                 string     `json:"source"`
	Active_ind             *string    `json:"active_ind"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Dls_event_sequence     *string    `json:"dls_event_sequence"`
	Dls_legal_subdivision  *float64   `json:"dls_legal_subdivision"`
	Dls_loc_exception      *string    `json:"dls_loc_exception"`
	Dls_meridian           *float64   `json:"dls_meridian"`
	Dls_meridian_direction *string    `json:"dls_meridian_direction"`
	Dls_range              *float64   `json:"dls_range"`
	Dls_range_modifier     *string    `json:"dls_range_modifier"`
	Dls_section            *float64   `json:"dls_section"`
	Dls_township           *float64   `json:"dls_township"`
	Dls_township_modifier  *string    `json:"dls_township_modifier"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Uwi                    *string    `json:"uwi"`
	Well_node_id           *string    `json:"well_node_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
