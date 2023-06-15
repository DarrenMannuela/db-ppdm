package dto

type Cont_account_proc struct {
	Contract_id          string   `json:"contract_id" default:"source"`
	Account_proc_type    string   `json:"account_proc_type" default:"source"`
	Account_proc_obs_no  int      `json:"account_proc_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Advance_percent      *float64 `json:"advance_percent" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Inventory_period     *float64 `json:"inventory_period" default:""`
	Inventory_period_uom *string  `json:"inventory_period_uom" default:""`
	Material_sale_limit  *float64 `json:"material_sale_limit" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Quorum_count         *int     `json:"quorum_count" default:""`
	Quorum_percent       *float64 `json:"quorum_percent" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Source_document_id   *string  `json:"source_document_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
