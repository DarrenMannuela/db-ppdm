package dto

import (
	"time"
)

type Seis_point struct {
	Seis_set_subtype    string     `json:"seis_set_subtype"`
	Seis_set_id         string     `json:"seis_set_id"`
	Seis_point_id       string     `json:"seis_point_id"`
	Acqtn_seq_no        *int       `json:"acqtn_seq_no"`
	Active_ind          *string    `json:"active_ind"`
	Bend_ind            *string    `json:"bend_ind"`
	Datum_elev          *float64   `json:"datum_elev"`
	Datum_elev_ouom     *string    `json:"datum_elev_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Elev                *float64   `json:"elev"`
	Elev_ouom           *string    `json:"elev_ouom"`
	End_ind             *string    `json:"end_ind"`
	Exception_ind       *string    `json:"exception_ind"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Flowing_hole_ind    *string    `json:"flowing_hole_ind"`
	Mapping_code        *string    `json:"mapping_code"`
	Measured_depth      *float64   `json:"measured_depth"`
	Measured_depth_ouom *string    `json:"measured_depth_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Reference_datum     *string    `json:"reference_datum"`
	Remark              *string    `json:"remark"`
	Seis_point_label    *string    `json:"seis_point_label"`
	Seis_point_lat      *float64   `json:"seis_point_lat"`
	Seis_point_long     *float64   `json:"seis_point_long"`
	Seis_point_no       *int       `json:"seis_point_no"`
	Seis_station_type   *string    `json:"seis_station_type"`
	Source              *string    `json:"source"`
	Spatial_seq_no      *int       `json:"spatial_seq_no"`
	X_coordinate        *float64   `json:"x_coordinate"`
	Y_coordinate        *float64   `json:"y_coordinate"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
