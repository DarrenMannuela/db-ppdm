package dto

type Rm_creator struct {
	Info_item_subtype   string  `json:"info_item_subtype" default:"source"`
	Information_item_id string  `json:"information_item_id" default:"source"`
	Creator_id          string  `json:"creator_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Creator_ba_id       *string `json:"creator_ba_id" default:""`
	Creator_ba_type     *string `json:"creator_ba_type" default:""`
	Creator_name        *string `json:"creator_name" default:""`
	Creator_type        *string `json:"creator_type" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	First_name          *string `json:"first_name" default:""`
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
