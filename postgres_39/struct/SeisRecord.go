package dto

import (
	"time"
)

type Seis_record struct {
	Seis_set_subtype       string     `json:"seis_set_subtype"`
	Seis_set_id            string     `json:"seis_set_id"`
	Record_id              string     `json:"record_id"`
	Active_ind             *string    `json:"active_ind"`
	Actual_acqtn_design_id *string    `json:"actual_acqtn_design_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Field_file_number      *string    `json:"field_file_number"`
	Information_item_id    *string    `json:"information_item_id"`
	Info_item_subtype      *string    `json:"info_item_subtype"`
	Logical_record_number  *string    `json:"logical_record_number"`
	Patch_id               *string    `json:"patch_id"`
	Patch_used_ind         *string    `json:"patch_used_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Rcrd_channel_count     *int       `json:"rcrd_channel_count"`
	Recording_remark       *string    `json:"recording_remark"`
	Record_number          *string    `json:"record_number"`
	Record_quality         *string    `json:"record_quality"`
	Record_type            *string    `json:"record_type"`
	Remark                 *string    `json:"remark"`
	Seis_shot_point_id     *string    `json:"seis_shot_point_id"`
	Source                 *string    `json:"source"`
	Tape_number            *string    `json:"tape_number"`
	Time_delay             *float64   `json:"time_delay"`
	Time_delay_ouom        *string    `json:"time_delay_ouom"`
	Uphole_time            *float64   `json:"uphole_time"`
	Uphole_time_ouom       *string    `json:"uphole_time_ouom"`
	Vessel_config_obs_no   *int       `json:"vessel_config_obs_no"`
	Vessel_id              *string    `json:"vessel_id"`
	Vessel_sf_subtype      *string    `json:"vessel_sf_subtype"`
	X_offset               *float64   `json:"x_offset"`
	Y_offset               *float64   `json:"y_offset"`
	Z_offset               *float64   `json:"z_offset"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
