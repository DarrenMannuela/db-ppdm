package dto

import (
	"time"
)

type Well_log_curve_frame struct {
	Uwi                string     `json:"uwi"`
	Well_log_id        string     `json:"well_log_id"`
	Well_log_source    string     `json:"well_log_source"`
	Frame_id           string     `json:"frame_id"`
	Active_ind         *string    `json:"active_ind"`
	Base_depth         *float64   `json:"base_depth"`
	Depth_ouom         *string    `json:"depth_ouom"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Frame_name         *string    `json:"frame_name"`
	Frame_spacing      *float64   `json:"frame_spacing"`
	Frame_spacing_ouom *string    `json:"frame_spacing_ouom"`
	Frame_spacing_uom  *string    `json:"frame_spacing_uom"`
	Gaps_ind           *string    `json:"gaps_ind"`
	Index_ouom         *string    `json:"index_ouom"`
	Index_uom          *string    `json:"index_uom"`
	Irregular_desc     *string    `json:"irregular_desc"`
	Irregular_ind      *string    `json:"irregular_ind"`
	Log_direction      *string    `json:"log_direction"`
	Max_index          *float64   `json:"max_index"`
	Min_index          *float64   `json:"min_index"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_index_type *string    `json:"primary_index_type"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Top_depth          *float64   `json:"top_depth"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
