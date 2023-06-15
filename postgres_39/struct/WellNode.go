package dto

import (
	"time"
)

type Well_node struct {
	Node_id               string     `json:"node_id"`
	Active_ind            *string    `json:"active_ind"`
	Coordinate_quality    *string    `json:"coordinate_quality"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Coord_system_id       *string    `json:"coord_system_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Latitude              *float64   `json:"latitude"`
	Legal_survey_type     *string    `json:"legal_survey_type"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Location_quality      *string    `json:"location_quality"`
	Location_type         *string    `json:"location_type"`
	Longitude             *float64   `json:"longitude"`
	Node_position         *string    `json:"node_position"`
	Original_obs_no       *int       `json:"original_obs_no"`
	Original_xy_uom       *string    `json:"original_xy_uom"`
	Platform_id           *string    `json:"platform_id"`
	Platform_sf_subtype   *string    `json:"platform_sf_subtype"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Remark                *string    `json:"remark"`
	Selected_obs_no       *int       `json:"selected_obs_no"`
	Source                *string    `json:"source"`
	Uwi                   *string    `json:"uwi"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
