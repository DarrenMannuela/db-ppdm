package dto

import (
	"time"
)

type Seis_point_summary struct {
	Seis_set_subtype    string     `json:"seis_set_subtype"`
	Seis_set_id         string     `json:"seis_set_id"`
	Seis_summary_obs_no int        `json:"seis_summary_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Area_size           *float64   `json:"area_size"`
	Area_size_ouom      *string    `json:"area_size_ouom"`
	Cdp_coverage        *float64   `json:"cdp_coverage"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	First_nline_no      *int       `json:"first_nline_no"`
	First_seis_point_id *string    `json:"first_seis_point_id"`
	First_xline_no      *int       `json:"first_xline_no"`
	Last_nline_no       *int       `json:"last_nline_no"`
	Last_seis_point_id  *string    `json:"last_seis_point_id"`
	Last_xline_no       *int       `json:"last_xline_no"`
	Line_length         *float64   `json:"line_length"`
	Line_length_ouom    *string    `json:"line_length_ouom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Seis_station_type   *string    `json:"seis_station_type"`
	Source              *string    `json:"source"`
	Summary_type        *string    `json:"summary_type"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
