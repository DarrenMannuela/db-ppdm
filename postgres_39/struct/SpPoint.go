package dto

import (
	"time"
)

type Sp_point struct {
	Point_id               string     `json:"point_id"`
	Active_ind             *string    `json:"active_ind"`
	Coord_acquisition_id   *string    `json:"coord_acquisition_id"`
	Coord_system_id        *string    `json:"coord_system_id"`
	Datum_elev             *float64   `json:"datum_elev"`
	Datum_elev_ouom        *string    `json:"datum_elev_ouom"`
	Effective_date         *time.Time `json:"effective_date"`
	Elevation              *float64   `json:"elevation"`
	Elevation_ouom         *string    `json:"elevation_ouom"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Latitude               *float64   `json:"latitude"`
	Local_coord_system_id  *string    `json:"local_coord_system_id"`
	Location_quality       *string    `json:"location_quality"`
	Location_type          *string    `json:"location_type"`
	Longitude              *float64   `json:"longitude"`
	Measured_depth         *float64   `json:"measured_depth"`
	Measured_depth_ouom    *string    `json:"measured_depth_ouom"`
	Point_label            *string    `json:"point_label"`
	Point_no               *float64   `json:"point_no"`
	Point_position         *string    `json:"point_position"`
	Point_seq_no           *int       `json:"point_seq_no"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Reference_datum        *string    `json:"reference_datum"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
