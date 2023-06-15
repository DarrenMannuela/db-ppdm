package dto

import (
	"time"
)

type Seis_bin_point_version struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Seis_set_id           string     `json:"seis_set_id"`
	Bin_grid_id           string     `json:"bin_grid_id"`
	Bin_grid_source       string     `json:"bin_grid_source"`
	Bin_point_id          string     `json:"bin_point_id"`
	Bin_point_obs_no      int        `json:"bin_point_obs_no"`
	Acquisition_id        *string    `json:"acquisition_id"`
	Active_ind            *string    `json:"active_ind"`
	Easting               *float64   `json:"easting"`
	Easting_ouom          *string    `json:"easting_ouom"`
	Effective_date        *time.Time `json:"effective_date"`
	Elev                  *float64   `json:"elev"`
	Elev_ouom             *string    `json:"elev_ouom"`
	Ew_direction          *string    `json:"ew_direction"`
	Ew_distance           *float64   `json:"ew_distance"`
	Ew_start_line         *string    `json:"ew_start_line"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Geog_coord_system_id  *string    `json:"geog_coord_system_id"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Map_coord_system_id   *string    `json:"map_coord_system_id"`
	Northing              *float64   `json:"northing"`
	Northing_ouom         *string    `json:"northing_ouom"`
	North_type            *string    `json:"north_type"`
	Ns_direction          *string    `json:"ns_direction"`
	Ns_distance           *float64   `json:"ns_distance"`
	Ns_start_line         *string    `json:"ns_start_line"`
	Polar_azimuth         *float64   `json:"polar_azimuth"`
	Polar_offset          *float64   `json:"polar_offset"`
	Polar_offset_ouom     *string    `json:"polar_offset_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Reference_loc         *string    `json:"reference_loc"`
	Remark                *string    `json:"remark"`
	Seis_point_label      *string    `json:"seis_point_label"`
	Seis_point_lat        *float64   `json:"seis_point_lat"`
	Seis_point_long       *float64   `json:"seis_point_long"`
	Seis_point_no         *int       `json:"seis_point_no"`
	Uwi                   *string    `json:"uwi"`
	Version_type          *string    `json:"version_type"`
	X_coordinate          *float64   `json:"x_coordinate"`
	Y_coordinate          *float64   `json:"y_coordinate"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
