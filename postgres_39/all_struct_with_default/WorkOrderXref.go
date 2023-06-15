package dto

type Work_order_xref struct {
	Work_order_id         string  `json:"work_order_id" default:"source"`
	Work_order_xref_id    string  `json:"work_order_xref_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Parent_work_order_id  *string `json:"parent_work_order_id" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Reference_id          *string `json:"reference_id" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Wo_xref_type          *string `json:"wo_xref_type" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
