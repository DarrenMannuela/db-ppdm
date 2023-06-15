package dto

import (
	"time"
)

type Seis_bin_outline struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Seis_set_id           string     `json:"seis_set_id"`
	Bin_grid_id           string     `json:"bin_grid_id"`
	Bin_grid_source       string     `json:"bin_grid_source"`
	Outline_seq_no        int        `json:"outline_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Bin_outline_type      *string    `json:"bin_outline_type"`
	Bin_point_id          *string    `json:"bin_point_id"`
	Coord_acquisition_id  *string    `json:"coord_acquisition_id"`
	Coord_system_id       *string    `json:"coord_system_id"`
	Easting               *float64   `json:"easting"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Nline_no              *int       `json:"nline_no"`
	Northing              *float64   `json:"northing"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Xline_no              *int       `json:"xline_no"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
