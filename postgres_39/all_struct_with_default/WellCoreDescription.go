package dto

type Well_core_description struct {
	Uwi                      string   `json:"uwi" default:"source"`
	Source                   string   `json:"source" default:"source"`
	Core_id                  string   `json:"core_id" default:"source"`
	Description_obs_no       int      `json:"description_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Base_depth               *float64 `json:"base_depth" default:""`
	Base_depth_ouom          *string  `json:"base_depth_ouom" default:""`
	Core_description_company *string  `json:"core_description_company" default:""`
	Description_date         *string  `json:"description_date" default:""`
	Dip_angle                *float64 `json:"dip_angle" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Interval_thickness       *float64 `json:"interval_thickness" default:""`
	Interval_thickness_ouom  *string  `json:"interval_thickness_ouom" default:""`
	Lithology_desc           *string  `json:"lithology_desc" default:""`
	Porosity_length          *float64 `json:"porosity_length" default:""`
	Porosity_length_ouom     *string  `json:"porosity_length_ouom" default:""`
	Porosity_quality         *string  `json:"porosity_quality" default:""`
	Porosity_type            *string  `json:"porosity_type" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Show_length              *float64 `json:"show_length" default:""`
	Show_length_ouom         *string  `json:"show_length_ouom" default:""`
	Show_type                *string  `json:"show_type" default:""`
	Top_depth                *float64 `json:"top_depth" default:""`
	Top_depth_ouom           *string  `json:"top_depth_ouom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
