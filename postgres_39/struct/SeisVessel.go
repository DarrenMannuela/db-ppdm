package dto

import (
	"time"
)

type Seis_vessel struct {
	Sf_subtype           string     `json:"sf_subtype"`
	Vessel_id            string     `json:"vessel_id"`
	Vessel_config_obs_no int        `json:"vessel_config_obs_no"`
	Acqtn_design_id      *string    `json:"acqtn_design_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Fath_azimuth         *float64   `json:"fath_azimuth"`
	Fath_offset          *float64   `json:"fath_offset"`
	Master_vessel_ind    *string    `json:"master_vessel_ind"`
	Offset_ouom          *string    `json:"offset_ouom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Reference_point      *string    `json:"reference_point"`
	Remark               *string    `json:"remark"`
	Shot_offset          *float64   `json:"shot_offset"`
	Slave_vessel_ind     *string    `json:"slave_vessel_ind"`
	Source               *string    `json:"source"`
	Streamer_far_offset  *float64   `json:"streamer_far_offset"`
	Streamer_near_offset *float64   `json:"streamer_near_offset"`
	Vessel_azimuth       *float64   `json:"vessel_azimuth"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
