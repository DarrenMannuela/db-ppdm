package dto

import (
	"time"
)

type Seis_velocity struct {
	Velocity_volume_id        string     `json:"velocity_volume_id"`
	Volume_point              string     `json:"volume_point"`
	Active_ind                *string    `json:"active_ind"`
	Bin_grid_id               *string    `json:"bin_grid_id"`
	Bin_point_id              *string    `json:"bin_point_id"`
	Bin_source                *string    `json:"bin_source"`
	Compute_method            *string    `json:"compute_method"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Latitude                  *float64   `json:"latitude"`
	Longitude                 *float64   `json:"longitude"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Receiver_md               *float64   `json:"receiver_md"`
	Receiver_md_ouom          *string    `json:"receiver_md_ouom"`
	Receiver_vert_depth       *float64   `json:"receiver_vert_depth"`
	Receiver_vert_depth_ouom  *string    `json:"receiver_vert_depth_ouom"`
	Remark                    *string    `json:"remark"`
	Seis_point_id             *string    `json:"seis_point_id"`
	Seis_set_id               *string    `json:"seis_set_id"`
	Seis_set_subtype          *string    `json:"seis_set_subtype"`
	Seis_time                 *float64   `json:"seis_time"`
	Seis_time_ouom            *string    `json:"seis_time_ouom"`
	Seis_well_set_id          *string    `json:"seis_well_set_id"`
	Seis_well_set_subtype     *string    `json:"seis_well_set_subtype"`
	Source                    *string    `json:"source"`
	Source_md                 *float64   `json:"source_md"`
	Source_md_ouom            *string    `json:"source_md_ouom"`
	Source_vert_depth         *float64   `json:"source_vert_depth"`
	Source_vert_depth_ouom    *string    `json:"source_vert_depth_ouom"`
	Velocity                  *float64   `json:"velocity"`
	Velocity_azimuth          *float64   `json:"velocity_azimuth"`
	Velocity_depth            *float64   `json:"velocity_depth"`
	Velocity_inclination      *float64   `json:"velocity_inclination"`
	Velocity_inclination_ouom *string    `json:"velocity_inclination_ouom"`
	Velocity_ouom             *string    `json:"velocity_ouom"`
	Velocity_type             *string    `json:"velocity_type"`
	Velocity_x                *float64   `json:"velocity_x"`
	Velocity_y                *float64   `json:"velocity_y"`
	Velocity_z                *float64   `json:"velocity_z"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
