package dto

type Fin_cost_summary struct {
	Cost_id              string   `json:"cost_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Ami_ind              *string  `json:"ami_ind" default:""`
	Confidential_ind     *string  `json:"confidential_ind" default:""`
	Cost_amount          *float64 `json:"cost_amount" default:""`
	Cost_per_size        *float64 `json:"cost_per_size" default:""`
	Cost_per_size_ouom   *string  `json:"cost_per_size_ouom" default:""`
	Cost_type            *string  `json:"cost_type" default:""`
	Currency_conversion  *float64 `json:"currency_conversion" default:""`
	Currency_ouom        *string  `json:"currency_ouom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Finance_component_id *string  `json:"finance_component_id" default:""`
	Finance_id           *string  `json:"finance_id" default:""`
	Paid_date            *string  `json:"paid_date" default:""`
	Parent_cost_id       *string  `json:"parent_cost_id" default:""`
	Percent_of_parent    *float64 `json:"percent_of_parent" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Report_cost_ind      *string  `json:"report_cost_ind" default:""`
	Source               *string  `json:"source" default:""`
	Submit_date          *string  `json:"submit_date" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
