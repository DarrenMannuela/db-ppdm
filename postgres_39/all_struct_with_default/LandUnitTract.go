package dto

type Land_unit_tract struct {
	Land_right_subtype     string   `json:"land_right_subtype" default:"source"`
	Land_right_id          string   `json:"land_right_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Land_tract_type        *string  `json:"land_tract_type" default:""`
	Land_unit_tract_name   *string  `json:"land_unit_tract_name" default:""`
	Land_unit_tract_number *string  `json:"land_unit_tract_number" default:""`
	Percent_crown          *float64 `json:"percent_crown" default:""`
	Percent_freehold       *float64 `json:"percent_freehold" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
