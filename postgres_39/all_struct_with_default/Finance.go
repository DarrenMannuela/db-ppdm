package dto

type Finance struct {
	Finance_id          string   `json:"finance_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Actual_cost         *float64 `json:"actual_cost" default:""`
	Authorized_by_ba_id *string  `json:"authorized_by_ba_id" default:""`
	Budget_cost         *float64 `json:"budget_cost" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Finance_type        *string  `json:"finance_type" default:""`
	Fin_status          *string  `json:"fin_status" default:""`
	Issue_date          *string  `json:"issue_date" default:""`
	Limit_amount        *float64 `json:"limit_amount" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Reference_number    *string  `json:"reference_number" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Tax_credit_code     *string  `json:"tax_credit_code" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
