package dto

import (
	"time"
)

type Seis_velocity_volume struct {
	Velocity_volume_id        string     `json:"velocity_volume_id"`
	Active_ind                *string    `json:"active_ind"`
	Bin_grid_id               *string    `json:"bin_grid_id"`
	Bin_grid_seis_set_id      *string    `json:"bin_grid_seis_set_id"`
	Bin_grid_seis_set_subtype *string    `json:"bin_grid_seis_set_subtype"`
	Bin_grid_source           *string    `json:"bin_grid_source"`
	Created_date              *time.Time `json:"created_date"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Picked_by                 *string    `json:"picked_by"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Seis_well_set_id          *string    `json:"seis_well_set_id"`
	Seis_well_set_subtype     *string    `json:"seis_well_set_subtype"`
	Source                    *string    `json:"source"`
	Velocity_dimension        *string    `json:"velocity_dimension"`
	Velocity_origin           *string    `json:"velocity_origin"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
