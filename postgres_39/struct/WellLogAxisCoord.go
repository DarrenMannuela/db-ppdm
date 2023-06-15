package dto

import (
	"time"
)

type Well_log_axis_coord struct {
	Uwi                string     `json:"uwi"`
	Curve_id           string     `json:"curve_id"`
	Axis_id            string     `json:"axis_id"`
	Coordinate_seq_no  int        `json:"coordinate_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Axis_value         *float64   `json:"axis_value"`
	Axis_value_ouom    *string    `json:"axis_value_ouom"`
	Axis_value_uom     *string    `json:"axis_value_uom"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Text_value         *string    `json:"text_value"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
