package dto

import (
	"time"
)

type Seis_streamer struct {
	Streamer_id          string     `json:"streamer_id"`
	Acqtn_design_id      *string    `json:"acqtn_design_id"`
	Active_ind           *string    `json:"active_ind"`
	Cable_make           *string    `json:"cable_make"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Sf_subtype           *string    `json:"sf_subtype"`
	Source               *string    `json:"source"`
	Streamer_length      *float64   `json:"streamer_length"`
	Streamer_length_ouom *string    `json:"streamer_length_ouom"`
	Streamer_position    *string    `json:"streamer_position"`
	Vessel_config_obs_no *int       `json:"vessel_config_obs_no"`
	Vessel_id            *string    `json:"vessel_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
