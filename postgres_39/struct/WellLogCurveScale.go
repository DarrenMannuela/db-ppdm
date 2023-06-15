package dto

import (
	"time"
)

type Well_log_curve_scale struct {
	Uwi                      string     `json:"uwi"`
	Curve_id                 string     `json:"curve_id"`
	Digital_curve_id         string     `json:"digital_curve_id"`
	Curve_scale_id           string     `json:"curve_scale_id"`
	Active_ind               *string    `json:"active_ind"`
	Backup_curve_scale       *string    `json:"backup_curve_scale"`
	Base_depth               *float64   `json:"base_depth"`
	Base_depth_ouom          *string    `json:"base_depth_ouom"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Left_scale_value         *float64   `json:"left_scale_value"`
	Matrix_lithology_setting *string    `json:"matrix_lithology_setting"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Right_scale_value        *float64   `json:"right_scale_value"`
	Scale_transform_type     *string    `json:"scale_transform_type"`
	Source                   *string    `json:"source"`
	Top_depth                *float64   `json:"top_depth"`
	Top_depth_ouom           *string    `json:"top_depth_ouom"`
	Track_num                *string    `json:"track_num"`
	Track_width              *float64   `json:"track_width"`
	Track_width_ouom         *string    `json:"track_width_ouom"`
	Track_width_uom          *string    `json:"track_width_uom"`
	Vertical_scale_ratio     *string    `json:"vertical_scale_ratio"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
