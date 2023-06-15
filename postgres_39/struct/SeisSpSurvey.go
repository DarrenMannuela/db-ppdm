package dto

import (
	"time"
)

type Seis_sp_survey struct {
	Seis_set_subtype      string     `json:"seis_set_subtype"`
	Seis_set_id           string     `json:"seis_set_id"`
	Seis_point_id         string     `json:"seis_point_id"`
	Seis_sp_survey_obs_no int        `json:"seis_sp_survey_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Azimuth               *float64   `json:"azimuth"`
	Azimuth_ouom          *string    `json:"azimuth_ouom"`
	Effective_date        *time.Time `json:"effective_date"`
	Ew_direction          *string    `json:"ew_direction"`
	Ew_distance           *float64   `json:"ew_distance"`
	Ew_distance_ouom      *string    `json:"ew_distance_ouom"`
	Ew_start_line         *string    `json:"ew_start_line"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Location_type         *string    `json:"location_type"`
	Monument_id           *string    `json:"monument_id"`
	Monument_sf_subtype   *string    `json:"monument_sf_subtype"`
	North_type            *string    `json:"north_type"`
	Ns_direction          *string    `json:"ns_direction"`
	Ns_distance           *float64   `json:"ns_distance"`
	Ns_distance_ouom      *string    `json:"ns_distance_ouom"`
	Ns_start_line         *string    `json:"ns_start_line"`
	Orientation           *string    `json:"orientation"`
	Point_offset          *float64   `json:"point_offset"`
	Point_offset_ouom     *string    `json:"point_offset_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Reference_loc         *string    `json:"reference_loc"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Surface_loc           *string    `json:"surface_loc"`
	Well_node_id          *string    `json:"well_node_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
