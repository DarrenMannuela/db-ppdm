package dto

type Rm_info_item_group struct {
	Group_info_item_subtype   string  `json:"group_info_item_subtype" default:"source"`
	Group_information_item_id string  `json:"group_information_item_id" default:"source"`
	Info_item_subtype         string  `json:"info_item_subtype" default:"source"`
	Information_item_id       string  `json:"information_item_id" default:"source"`
	Group_obs_no              int     `json:"group_obs_no" default:"1"`
	Active_ind                *string `json:"active_ind" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Group_seq_no              *int    `json:"group_seq_no" default:""`
	Group_type                *string `json:"group_type" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Remark                    *string `json:"remark" default:""`
	Source                    *string `json:"source" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
