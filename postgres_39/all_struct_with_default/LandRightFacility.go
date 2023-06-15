package dto

type Land_right_facility struct {
	Land_right_subtype string  `json:"land_right_subtype" default:"source"`
	Land_right_id      string  `json:"land_right_id" default:"source"`
	Facility_id        string  `json:"facility_id" default:"source"`
	Facility_type      string  `json:"facility_type" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Unit_operated_ind  *string `json:"unit_operated_ind" default:""`
	Xref_type          *string `json:"xref_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
