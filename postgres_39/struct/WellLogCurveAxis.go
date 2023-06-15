package dto

import (
	"time"
)

type Well_log_curve_axis struct {
	Uwi                       string     `json:"uwi"`
	Curve_id                  string     `json:"curve_id"`
	Axis_id                   string     `json:"axis_id"`
	Active_ind                *string    `json:"active_ind"`
	Axis_ouom                 *string    `json:"axis_ouom"`
	Axis_seq_no               *int       `json:"axis_seq_no"`
	Axis_uom                  *string    `json:"axis_uom"`
	Dimension_count           *int       `json:"dimension_count"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Reported_axis_name        *string    `json:"reported_axis_name"`
	Reported_axis_object_name *string    `json:"reported_axis_object_name"`
	Source                    *string    `json:"source"`
	Spacing                   *float64   `json:"spacing"`
	Spacing_ouom              *string    `json:"spacing_ouom"`
	Spacing_uom               *string    `json:"spacing_uom"`
	Value_count               *int       `json:"value_count"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
