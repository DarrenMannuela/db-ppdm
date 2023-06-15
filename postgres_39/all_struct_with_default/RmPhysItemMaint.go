package dto

type Rm_phys_item_maint struct {
	Physical_item_id   string  `json:"physical_item_id" default:"source"`
	Maint_id           string  `json:"maint_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Actual_date        *string `json:"actual_date" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Maintain_process   *string `json:"maintain_process" default:""`
	Maint_ba_id        *string `json:"maint_ba_id" default:""`
	Maint_by_name      *string `json:"maint_by_name" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Scheduled_date     *string `json:"scheduled_date" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
