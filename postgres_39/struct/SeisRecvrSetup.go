package dto

import (
	"time"
)

type Seis_recvr_setup struct {
	Acqtn_design_id          string     `json:"acqtn_design_id"`
	Rcvr_setup_id            string     `json:"rcvr_setup_id"`
	Active_ind               *string    `json:"active_ind"`
	Avg_feathering_angle     *float64   `json:"avg_feathering_angle"`
	Avg_streamer_depth       *float64   `json:"avg_streamer_depth"`
	Avg_streamer_depth_ouom  *string    `json:"avg_streamer_depth_ouom"`
	Base_freq                *float64   `json:"base_freq"`
	Depth_controller         *string    `json:"depth_controller"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Fixed_ind                *string    `json:"fixed_ind"`
	Group_spacing            *float64   `json:"group_spacing"`
	Group_spacing_ouom       *string    `json:"group_spacing_ouom"`
	Inline_offset            *float64   `json:"inline_offset"`
	Inline_offset_direction  *string    `json:"inline_offset_direction"`
	Offline_offset           *float64   `json:"offline_offset"`
	Offline_offset_direction *string    `json:"offline_offset_direction"`
	Offset_ouom              *string    `json:"offset_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Rcvr_array_type          *string    `json:"rcvr_array_type"`
	Rcvr_make                *string    `json:"rcvr_make"`
	Rcvr_phone_count         *int       `json:"rcvr_phone_count"`
	Rcvr_spacing             *float64   `json:"rcvr_spacing"`
	Rcvr_spacing_ouom        *string    `json:"rcvr_spacing_ouom"`
	Rcvr_type                *string    `json:"rcvr_type"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Spread_description       *string    `json:"spread_description"`
	Spread_description_ouom  *string    `json:"spread_description_ouom"`
	Streamer_count           *int       `json:"streamer_count"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
