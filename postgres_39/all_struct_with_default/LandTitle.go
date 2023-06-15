package dto

type Land_title struct {
	Land_right_subtype  string  `json:"land_right_subtype" default:"source"`
	Land_right_id       string  `json:"land_right_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Certified_desc      *string `json:"certified_desc" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Registration_date   *string `json:"registration_date" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Title_change_reason *string `json:"title_change_reason" default:""`
	Title_holder        *string `json:"title_holder" default:""`
	Title_number        *string `json:"title_number" default:""`
	Title_reference_num *string `json:"title_reference_num" default:""`
	Title_type          *string `json:"title_type" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
