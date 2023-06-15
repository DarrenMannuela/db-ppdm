package dto

import (
	"time"
)

type Seis_bin_point struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Seis_set_id           string     `json:"seis_set_id"`
	Bin_grid_id           string     `json:"bin_grid_id"`
	Bin_grid_source       string     `json:"bin_grid_source"`
	Bin_point_id          string     `json:"bin_point_id"`
	Active_ind            *string    `json:"active_ind"`
	Easting               *float64   `json:"easting"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Nline_no              *int       `json:"nline_no"`
	Northing              *float64   `json:"northing"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Velocity_analysis_ind *string    `json:"velocity_analysis_ind"`
	Xline_no              *int       `json:"xline_no"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
