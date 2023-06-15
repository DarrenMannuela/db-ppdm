package dto

import (
	"time"
)

type Seis_streamer_build struct {
	Streamer_id           string     `json:"streamer_id"`
	Streamer_component_id string     `json:"streamer_component_id"`
	Active_ind            *string    `json:"active_ind"`
	Component_type        *string    `json:"component_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Position_ouom         *string    `json:"position_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Streamer_position     *float64   `json:"streamer_position"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
