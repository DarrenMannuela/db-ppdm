package dto

import (
	"time"
)

type Spatial_description struct {
	Spatial_description_id string     `json:"spatial_description_id"`
	Spatial_obs_no         int        `json:"spatial_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Carter_ind             *string    `json:"carter_ind"`
	Congress_ind           *string    `json:"congress_ind"`
	Coord_acquisition_id   *string    `json:"coord_acquisition_id"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Dls_ind                *string    `json:"dls_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Fps_ind                *string    `json:"fps_ind"`
	Geodetic_ind           *string    `json:"geodetic_ind"`
	Gross_size             *float64   `json:"gross_size"`
	Gross_size_ouom        *string    `json:"gross_size_ouom"`
	Inactivation_date      *time.Time `json:"inactivation_date"`
	Libya_ind              *string    `json:"libya_ind"`
	Line_ind               *string    `json:"line_ind"`
	Line_version_ind       *string    `json:"line_version_ind"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Max_latitude           *float64   `json:"max_latitude"`
	Max_longitude          *float64   `json:"max_longitude"`
	Mineral_zone_ind       *string    `json:"mineral_zone_ind"`
	Min_latitude           *float64   `json:"min_latitude"`
	Min_longitude          *float64   `json:"min_longitude"`
	Ne_loc_ind             *string    `json:"ne_loc_ind"`
	North_sea_ind          *string    `json:"north_sea_ind"`
	Nts_ind                *string    `json:"nts_ind"`
	Offshore_ind           *string    `json:"offshore_ind"`
	Ohio_ind               *string    `json:"ohio_ind"`
	Pbl_ind                *string    `json:"pbl_ind"`
	Point_ind              *string    `json:"point_ind"`
	Point_version_ind      *string    `json:"point_version_ind"`
	Polygon_ind            *string    `json:"polygon_ind"`
	Polygon_version_ind    *string    `json:"polygon_version_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_num          *string    `json:"reference_num"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_desc_text_ind  *string    `json:"spatial_desc_text_ind"`
	Spatial_desc_type      *string    `json:"spatial_desc_type"`
	Texas_ind              *string    `json:"texas_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
