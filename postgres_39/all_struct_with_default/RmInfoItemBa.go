package dto

type Rm_info_item_ba struct {
	Info_item_subtype   string  `json:"info_item_subtype" default:"source"`
	Information_item_id string  `json:"information_item_id" default:"source"`
	Contact_id          string  `json:"contact_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Contact_ba_id       *string `json:"contact_ba_id" default:""`
	Contact_ba_type     *string `json:"contact_ba_type" default:""`
	Contact_full_name   *string `json:"contact_full_name" default:""`
	Contact_type        *string `json:"contact_type" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	First_name          *string `json:"first_name" default:""`
	Instruction         *string `json:"instruction" default:""`
	Last_name           *string `json:"last_name" default:""`
	Middle_name         *string `json:"middle_name" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
