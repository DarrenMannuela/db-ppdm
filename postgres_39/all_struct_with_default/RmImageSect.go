package dto

type Rm_image_sect struct {
	Physical_item_id          string   `json:"physical_item_id" default:"source"`
	Image_section_id          string   `json:"image_section_id" default:"source"`
	Active_ind                *string  `json:"active_ind" default:""`
	Base_depth                *float64 `json:"base_depth" default:""`
	Base_depth_ouom           *string  `json:"base_depth_ouom" default:""`
	Calibrated_by_ba_id       *string  `json:"calibrated_by_ba_id" default:""`
	Calibrate_application     *string  `json:"calibrate_application" default:""`
	Calibrate_method          *string  `json:"calibrate_method" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Image_section_type        *string  `json:"image_section_type" default:""`
	Index_type                *string  `json:"index_type" default:""`
	Log_matrix_type           *string  `json:"log_matrix_type" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Scale_depth_interval      *float64 `json:"scale_depth_interval" default:""`
	Scale_depth_interval_ouom *string  `json:"scale_depth_interval_ouom" default:""`
	Scale_depth_interval_uom  *string  `json:"scale_depth_interval_uom" default:""`
	Scale_length              *float64 `json:"scale_length" default:""`
	Scale_length_ouom         *string  `json:"scale_length_ouom" default:""`
	Scale_length_uom          *string  `json:"scale_length_uom" default:""`
	Scale_ratio               *string  `json:"scale_ratio" default:""`
	Section_desc              *string  `json:"section_desc" default:""`
	Source                    *string  `json:"source" default:""`
	Top_depth                 *float64 `json:"top_depth" default:""`
	Top_depth_ouom            *string  `json:"top_depth_ouom" default:""`
	Well_log_class_id         *string  `json:"well_log_class_id" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
