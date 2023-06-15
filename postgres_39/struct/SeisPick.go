package dto

import (
	"time"
)

type Seis_pick struct {
	Seis_set_subtype       string     `json:"seis_set_subtype"`
	Interp_set_id          string     `json:"interp_set_id"`
	Surface_id             string     `json:"surface_id"`
	Seis_pick_id           string     `json:"seis_pick_id"`
	Active_ind             *string    `json:"active_ind"`
	Bin_grid_id            *string    `json:"bin_grid_id"`
	Bin_point_id           *string    `json:"bin_point_id"`
	Bin_seis_set_id        *string    `json:"bin_seis_set_id"`
	Bin_seis_set_subtype   *string    `json:"bin_seis_set_subtype"`
	Bin_source             *string    `json:"bin_source"`
	Date_format_desc       *time.Time `json:"date_format_desc"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Interp_type            *string    `json:"interp_type"`
	Pick_depth             *float64   `json:"pick_depth"`
	Pick_depth_ouom        *string    `json:"pick_depth_ouom"`
	Pick_description       *string    `json:"pick_description"`
	Pick_method            *string    `json:"pick_method"`
	Pick_qualifier         *string    `json:"pick_qualifier"`
	Pick_qualif_reason     *string    `json:"pick_qualif_reason"`
	Pick_quality           *string    `json:"pick_quality"`
	Pick_version_type      *string    `json:"pick_version_type"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_pick_ind     *string    `json:"preferred_pick_ind"`
	Remark                 *string    `json:"remark"`
	Seis_pick_value        *float64   `json:"seis_pick_value"`
	Seis_pick_value_ouom   *string    `json:"seis_pick_value_ouom"`
	Seis_pick_value_uom    *string    `json:"seis_pick_value_uom"`
	Seis_point_id          *string    `json:"seis_point_id"`
	Seis_point_set_id      *string    `json:"seis_point_set_id"`
	Seis_point_set_subtype *string    `json:"seis_point_set_subtype"`
	Source                 *string    `json:"source"`
	Trace_id               *string    `json:"trace_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
