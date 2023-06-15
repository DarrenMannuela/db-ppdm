package dto

import (
	"time"
)

type Seis_channel struct {
	Seis_set_subtype       string     `json:"seis_set_subtype"`
	Seis_set_id            string     `json:"seis_set_id"`
	Record_id              string     `json:"record_id"`
	Channel_id             string     `json:"channel_id"`
	Active_ind             *string    `json:"active_ind"`
	Channel_num            *string    `json:"channel_num"`
	Channel_type           *string    `json:"channel_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	File_num               *string    `json:"file_num"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Seis_receiver_point_id *string    `json:"seis_receiver_point_id"`
	Source                 *string    `json:"source"`
	Streamer_id            *string    `json:"streamer_id"`
	Vessel_config_obs_no   *int       `json:"vessel_config_obs_no"`
	Vessel_id              *string    `json:"vessel_id"`
	Vessel_sf_subtype      *string    `json:"vessel_sf_subtype"`
	X_offset               *float64   `json:"x_offset"`
	X_offset_ouom          *string    `json:"x_offset_ouom"`
	Y_offset               *float64   `json:"y_offset"`
	Y_offset_ouom          *string    `json:"y_offset_ouom"`
	Z_offset               *float64   `json:"z_offset"`
	Z_offset_ouom          *string    `json:"z_offset_ouom"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
