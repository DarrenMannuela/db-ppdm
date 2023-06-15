package dto

import (
	"time"
)

type Legal_fps_loc struct {
	Legal_fps_loc_id    string     `json:"legal_fps_loc_id"`
	Location_type       string     `json:"location_type"`
	Source              string     `json:"source"`
	Active_ind          *string    `json:"active_ind"`
	Coord_system_id     *string    `json:"coord_system_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Fps_event_sequence  *string    `json:"fps_event_sequence"`
	Fps_loc_exception   *string    `json:"fps_loc_exception"`
	Grid                *string    `json:"grid"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Reference_latitude  *float64   `json:"reference_latitude"`
	Reference_longitude *float64   `json:"reference_longitude"`
	Remark              *string    `json:"remark"`
	Section             *int       `json:"section"`
	Unit                *string    `json:"unit"`
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
