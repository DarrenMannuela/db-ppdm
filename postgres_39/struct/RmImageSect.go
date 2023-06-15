package dto

import (
	"time"
)

type Rm_image_sect struct {
	Physical_item_id          string     `json:"physical_item_id"`
	Image_section_id          string     `json:"image_section_id"`
	Active_ind                *string    `json:"active_ind"`
	Base_depth                *float64   `json:"base_depth"`
	Base_depth_ouom           *string    `json:"base_depth_ouom"`
	Calibrated_by_ba_id       *string    `json:"calibrated_by_ba_id"`
	Calibrate_application     *string    `json:"calibrate_application"`
	Calibrate_method          *string    `json:"calibrate_method"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Image_section_type        *string    `json:"image_section_type"`
	Index_type                *string    `json:"index_type"`
	Log_matrix_type           *string    `json:"log_matrix_type"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Scale_depth_interval      *float64   `json:"scale_depth_interval"`
	Scale_depth_interval_ouom *string    `json:"scale_depth_interval_ouom"`
	Scale_depth_interval_uom  *string    `json:"scale_depth_interval_uom"`
	Scale_length              *float64   `json:"scale_length"`
	Scale_length_ouom         *string    `json:"scale_length_ouom"`
	Scale_length_uom          *string    `json:"scale_length_uom"`
	Scale_ratio               *string    `json:"scale_ratio"`
	Section_desc              *string    `json:"section_desc"`
	Source                    *string    `json:"source"`
	Top_depth                 *float64   `json:"top_depth"`
	Top_depth_ouom            *string    `json:"top_depth_ouom"`
	Well_log_class_id         *string    `json:"well_log_class_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
