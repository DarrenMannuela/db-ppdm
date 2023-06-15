package dto

type Cont_allow_expense struct {
	Contract_id             string   `json:"contract_id" default:"source"`
	Allow_expense_type      string   `json:"allow_expense_type" default:"source"`
	Allow_expense_obs_no    int      `json:"allow_expense_obs_no" default:"1"`
	Account_proc_obs_no     *int     `json:"account_proc_obs_no" default:""`
	Account_proc_type       *string  `json:"account_proc_type" default:""`
	Active_ind              *string  `json:"active_ind" default:""`
	Allow_percent           *float64 `json:"allow_percent" default:""`
	Allow_percent_qualifier *string  `json:"allow_percent_qualifier" default:""`
	Currency_conversion     *float64 `json:"currency_conversion" default:""`
	Currency_ouom           *string  `json:"currency_ouom" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expense_frequency       *float64 `json:"expense_frequency" default:""`
	Expense_frequency_uom   *string  `json:"expense_frequency_uom" default:""`
	Expense_rate            *float64 `json:"expense_rate" default:""`
	Expense_rate_qualifier  *string  `json:"expense_rate_qualifier" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ineligible_ind          *string  `json:"ineligible_ind" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
