package dto

import (
	"time"
)

type Seis_bin_grid struct {
	Seis_set_subtype       string     `json:"seis_set_subtype"`
	Seis_set_id            string     `json:"seis_set_id"`
	Bin_grid_id            string     `json:"bin_grid_id"`
	Source                 string     `json:"source"`
	Active_ind             *string    `json:"active_ind"`
	Angle_rotation         *float64   `json:"angle_rotation"`
	Bin_grid_type          *string    `json:"bin_grid_type"`
	Bin_grid_version       *float64   `json:"bin_grid_version"`
	Bin_method             *string    `json:"bin_method"`
	Bin_point_ouom         *string    `json:"bin_point_ouom"`
	Coord_acquisition_id   *string    `json:"coord_acquisition_id"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Corner1_lat            *float64   `json:"corner_1_lat"`
	Corner1_long           *float64   `json:"corner_1_long"`
	Corner2_lat            *float64   `json:"corner_2_lat"`
	Corner2_long           *float64   `json:"corner_2_long"`
	Corner3_lat            *float64   `json:"corner_3_lat"`
	Corner3_long           *float64   `json:"corner_3_long"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Fold_coverage          *float64   `json:"fold_coverage"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Nline_azimuth          *float64   `json:"nline_azimuth"`
	Nline_count            *int       `json:"nline_count"`
	Nline_max_no           *int       `json:"nline_max_no"`
	Nline_min_no           *int       `json:"nline_min_no"`
	Nline_spacing          *float64   `json:"nline_spacing"`
	North_type             *string    `json:"north_type"`
	Point_origin_easting   *float64   `json:"point_origin_easting"`
	Point_origin_latitude  *float64   `json:"point_origin_latitude"`
	Point_origin_longitude *float64   `json:"point_origin_longitude"`
	Point_origin_northing  *float64   `json:"point_origin_northing"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Xline_azimuth          *float64   `json:"xline_azimuth"`
	Xline_count            *int       `json:"xline_count"`
	Xline_max_no           *int       `json:"xline_max_no"`
	Xline_min_no           *int       `json:"xline_min_no"`
	Xline_spacing          *float64   `json:"xline_spacing"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
