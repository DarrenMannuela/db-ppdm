package dto

type Reserve_entity struct {
	Resent_id            string  `json:"resent_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Created_by_ba_id     *string `json:"created_by_ba_id" default:""`
	Description          *string `json:"description" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Group_type           *string `json:"group_type" default:""`
	Last_approve_ba_id   *string `json:"last_approve_ba_id" default:""`
	Last_update_ba_id    *string `json:"last_update_ba_id" default:""`
	Last_update_date     *string `json:"last_update_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Primary_product_type *string `json:"primary_product_type" default:""`
	Remark               *string `json:"remark" default:""`
	Report_ind           *string `json:"report_ind" default:""`
	Resent_long_name     *string `json:"resent_long_name" default:""`
	Resent_short_name    *string `json:"resent_short_name" default:""`
	Source               *string `json:"source" default:""`
	Update_schedule      *string `json:"update_schedule" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
