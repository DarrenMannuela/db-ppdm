package dto

import (
	"time"
)

type Strat_field_node struct {
	Field_station_id      string     `json:"field_station_id"`
	Node_id               string     `json:"node_id"`
	Source                string     `json:"source"`
	Active_ind            *string    `json:"active_ind"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Coord_system_id       *string    `json:"coord_system_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Elev                  *float64   `json:"elev"`
	Elev_ouom             *string    `json:"elev_ouom"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Latitude              *float64   `json:"latitude"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Longitude             *float64   `json:"longitude"`
	Node_loc_type         *string    `json:"node_loc_type"`
	Original_obs_no       *int       `json:"original_obs_no"`
	Original_xy_uom       *string    `json:"original_xy_uom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Selected_obs_no       *int       `json:"selected_obs_no"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
