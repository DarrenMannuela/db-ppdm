package dto

import (
	"time"
)

type Sp_point_version struct {
	Point_id              string     `json:"point_id"`
	Source                string     `json:"source"`
	Version_obs_no        int        `json:"version_obs_no"`
	Acquisition_id        *string    `json:"acquisition_id"`
	Active_ind            *string    `json:"active_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Easting               *float64   `json:"easting"`
	Easting_ouom          *string    `json:"easting_ouom"`
	Effective_date        *time.Time `json:"effective_date"`
	Elev                  *float64   `json:"elev"`
	Elev_ouom             *string    `json:"elev_ouom"`
	Ew_direction          *string    `json:"ew_direction"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Geog_coord_system_id  *string    `json:"geog_coord_system_id"`
	Latitude              *float64   `json:"latitude"`
	Legal_survey_type     *string    `json:"legal_survey_type"`
	Local_coord_system_id *string    `json:"local_coord_system_id"`
	Location_qualifier    *string    `json:"location_qualifier"`
	Location_quality      *string    `json:"location_quality"`
	Longitude             *float64   `json:"longitude"`
	Map_coord_system_id   *string    `json:"map_coord_system_id"`
	Md                    *float64   `json:"md"`
	Md_ouom               *string    `json:"md_ouom"`
	Monument_id           *string    `json:"monument_id"`
	Monument_sf_subtype   *string    `json:"monument_sf_subtype"`
	Node_position         *string    `json:"node_position"`
	Northing              *float64   `json:"northing"`
	Northing_ouom         *string    `json:"northing_ouom"`
	North_type            *string    `json:"north_type"`
	Ns_direction          *string    `json:"ns_direction"`
	Original_ind          *string    `json:"original_ind"`
	Original_ref_num      *string    `json:"original_ref_num"`
	Polar_azimuth         *float64   `json:"polar_azimuth"`
	Polar_offset          *float64   `json:"polar_offset"`
	Polar_offset_ouom     *string    `json:"polar_offset_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Remark                *string    `json:"remark"`
	Reported_tvd          *float64   `json:"reported_tvd"`
	Reported_tvd_ouom     *string    `json:"reported_tvd_ouom"`
	Version_type          *string    `json:"version_type"`
	X_coordinate          *float64   `json:"x_coordinate"`
	X_offset              *float64   `json:"x_offset"`
	X_offset_ouom         *string    `json:"x_offset_ouom"`
	Y_coordinate          *float64   `json:"y_coordinate"`
	Y_offset              *float64   `json:"y_offset"`
	Y_offset_ouom         *string    `json:"y_offset_ouom"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
