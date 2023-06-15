package dto

import (
	"time"
)

type Well_log_dgtz_curve struct {
	Uwi                     string     `json:"uwi"`
	Curve_id                string     `json:"curve_id"`
	Digital_curve_id        string     `json:"digital_curve_id"`
	Active_ind              *string    `json:"active_ind"`
	Base_depth              *float64   `json:"base_depth"`
	Base_depth_ouom         *string    `json:"base_depth_ouom"`
	Curve_quality           *string    `json:"curve_quality"`
	Depth_correction_method *string    `json:"depth_correction_method"`
	Depth_increment         *float64   `json:"depth_increment"`
	Depth_increment_ouom    *string    `json:"depth_increment_ouom"`
	Digitized_date          *time.Time `json:"digitized_date"`
	Digitizer               *string    `json:"digitizer"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Null_value              *float64   `json:"null_value"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Top_depth               *float64   `json:"top_depth"`
	Top_depth_ouom          *string    `json:"top_depth_ouom"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
