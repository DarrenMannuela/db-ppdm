package dto

type Rm_info_item_maint struct {
	Info_item_subtype   string   `json:"info_item_subtype" default:"source"`
	Information_item_id string   `json:"information_item_id" default:"source"`
	Maintain_id         string   `json:"maintain_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Instruction         *string  `json:"instruction" default:""`
	Maint_ba_id         *string  `json:"maint_ba_id" default:""`
	Maint_complete_date *string  `json:"maint_complete_date" default:""`
	Maint_due_date      *string  `json:"maint_due_date" default:""`
	Maint_period        *float64 `json:"maint_period" default:""`
	Maint_period_type   *string  `json:"maint_period_type" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Scheduled_ind       *string  `json:"scheduled_ind" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
