package dto

import (
	"time"
)

type Legal_geodetic_loc struct {
	Legal_geodetic_loc_id   string     `json:"legal_geodetic_loc_id"`
	Location_type           string     `json:"location_type"`
	Source                  string     `json:"source"`
	Active_ind              *string    `json:"active_ind"`
	Coord_system_id         *string    `json:"coord_system_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Geodetic_event_sequence *string    `json:"geodetic_event_sequence"`
	Geodetic_loc_exception  *string    `json:"geodetic_loc_exception"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
	Node_id                 *string    `json:"node_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Uwi                     *string    `json:"uwi"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
