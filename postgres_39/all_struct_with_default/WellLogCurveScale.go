package dto

type Well_log_curve_scale struct {
	Uwi                      string   `json:"uwi" default:"source"`
	Curve_id                 string   `json:"curve_id" default:"source"`
	Digital_curve_id         string   `json:"digital_curve_id" default:"source"`
	Curve_scale_id           string   `json:"curve_scale_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Backup_curve_scale       *string  `json:"backup_curve_scale" default:""`
	Base_depth               *float64 `json:"base_depth" default:""`
	Base_depth_ouom          *string  `json:"base_depth_ouom" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Left_scale_value         *float64 `json:"left_scale_value" default:""`
	Matrix_lithology_setting *string  `json:"matrix_lithology_setting" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Right_scale_value        *float64 `json:"right_scale_value" default:""`
	Scale_transform_type     *string  `json:"scale_transform_type" default:""`
	Source                   *string  `json:"source" default:""`
	Top_depth                *float64 `json:"top_depth" default:""`
	Top_depth_ouom           *string  `json:"top_depth_ouom" default:""`
	Track_num                *string  `json:"track_num" default:""`
	Track_width              *float64 `json:"track_width" default:""`
	Track_width_ouom         *string  `json:"track_width_ouom" default:""`
	Track_width_uom          *string  `json:"track_width_uom" default:""`
	Vertical_scale_ratio     *string  `json:"vertical_scale_ratio" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
