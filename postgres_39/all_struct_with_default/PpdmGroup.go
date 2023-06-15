package dto

type Ppdm_group struct {
	Group_id               string  `json:"group_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Default_unit_system_id *string `json:"default_unit_system_id" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Group_name             *string `json:"group_name" default:""`
	Group_type             *string `json:"group_type" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Report_heading1        *string `json:"report_heading_1" default:""`
	Report_heading2        *string `json:"report_heading_2" default:""`
	Screen_heading1        *string `json:"screen_heading_1" default:""`
	Screen_heading2        *string `json:"screen_heading_2" default:""`
	Source                 *string `json:"source" default:""`
	Sw_application_id      *string `json:"sw_application_id" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
