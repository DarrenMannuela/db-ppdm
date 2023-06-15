package dto

type Well_log_curve_axis struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Curve_id                  string   `json:"curve_id" default:"source"`
	Axis_id                   string   `json:"axis_id" default:"source"`
	Active_ind                *string  `json:"active_ind" default:""`
	Axis_ouom                 *string  `json:"axis_ouom" default:""`
	Axis_seq_no               *int     `json:"axis_seq_no" default:""`
	Axis_uom                  *string  `json:"axis_uom" default:""`
	Dimension_count           *int     `json:"dimension_count" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Reported_axis_name        *string  `json:"reported_axis_name" default:""`
	Reported_axis_object_name *string  `json:"reported_axis_object_name" default:""`
	Source                    *string  `json:"source" default:""`
	Spacing                   *float64 `json:"spacing" default:""`
	Spacing_ouom              *string  `json:"spacing_ouom" default:""`
	Spacing_uom               *string  `json:"spacing_uom" default:""`
	Value_count               *int     `json:"value_count" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
