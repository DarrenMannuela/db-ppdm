package dto

type Rm_phys_item_group struct {
	Group_physical_item_id string  `json:"group_physical_item_id" default:"source"`
	Physical_item_id       string  `json:"physical_item_id" default:"source"`
	Group_obs_no           int     `json:"group_obs_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Group_type             *string `json:"group_type" default:""`
	Order_seq_no           *int    `json:"order_seq_no" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
