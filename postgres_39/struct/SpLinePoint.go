package dto

import (
	"time"
)

type Sp_line_point struct {
	Line_id            string     `json:"line_id"`
	Point_seq_no       int        `json:"point_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Bend_ind           *string    `json:"bend_ind"`
	Depth              *float64   `json:"depth"`
	Depth_ouom         *string    `json:"depth_ouom"`
	Effective_date     *time.Time `json:"effective_date"`
	Elevation          *float64   `json:"elevation"`
	Elevation_ouom     *string    `json:"elevation_ouom"`
	Expiry_date        *time.Time `json:"expiry_date"`
	First_point_ind    *string    `json:"first_point_ind"`
	Last_point_ind     *string    `json:"last_point_ind"`
	Latitude           *float64   `json:"latitude"`
	Location_quality   *string    `json:"location_quality"`
	Longitude          *float64   `json:"longitude"`
	Point_label        *string    `json:"point_label"`
	Point_no           *float64   `json:"point_no"`
	Point_position     *string    `json:"point_position"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
