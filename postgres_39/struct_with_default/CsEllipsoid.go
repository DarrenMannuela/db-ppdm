package dto

type Cs_ellipsoid struct {
	Ellipsoid_id             string   `json:"ellipsoid_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Axis_ouom                *string  `json:"axis_ouom" default:""`
	Eccentricity             *float64 `json:"eccentricity" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Ellipsoid_name           *string  `json:"ellipsoid_name" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Inverse_flattening       *float64 `json:"inverse_flattening" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Semi_major_axis          *float64 `json:"semi_major_axis" default:""`
	Semi_major_axis_accuracy *int     `json:"semi_major_axis_accuracy" default:""`
	Semi_minor_axis          *float64 `json:"semi_minor_axis" default:""`
	Semi_minor_axis_accuracy *int     `json:"semi_minor_axis_accuracy" default:""`
	Source                   *string  `json:"source" default:""`
	Source_document_id       *string  `json:"source_document_id" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
