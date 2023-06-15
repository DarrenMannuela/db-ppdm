package dto

type Well_log_dgtz_curve struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Curve_id                string   `json:"curve_id" default:"source"`
	Digital_curve_id        string   `json:"digital_curve_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Base_depth              *float64 `json:"base_depth" default:""`
	Base_depth_ouom         *string  `json:"base_depth_ouom" default:""`
	Curve_quality           *string  `json:"curve_quality" default:""`
	Depth_correction_method *string  `json:"depth_correction_method" default:""`
	Depth_increment         *float64 `json:"depth_increment" default:""`
	Depth_increment_ouom    *string  `json:"depth_increment_ouom" default:""`
	Digitized_date          *string  `json:"digitized_date" default:""`
	Digitizer               *string  `json:"digitizer" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Null_value              *float64 `json:"null_value" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Top_depth               *float64 `json:"top_depth" default:""`
	Top_depth_ouom          *string  `json:"top_depth_ouom" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
