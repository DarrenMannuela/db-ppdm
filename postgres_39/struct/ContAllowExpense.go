package dto

import (
	"time"
)

type Cont_allow_expense struct {
	Contract_id             string     `json:"contract_id"`
	Allow_expense_type      string     `json:"allow_expense_type"`
	Allow_expense_obs_no    int        `json:"allow_expense_obs_no"`
	Account_proc_obs_no     *int       `json:"account_proc_obs_no"`
	Account_proc_type       *string    `json:"account_proc_type"`
	Active_ind              *string    `json:"active_ind"`
	Allow_percent           *float64   `json:"allow_percent"`
	Allow_percent_qualifier *string    `json:"allow_percent_qualifier"`
	Currency_conversion     *float64   `json:"currency_conversion"`
	Currency_ouom           *string    `json:"currency_ouom"`
	Effective_date          *time.Time `json:"effective_date"`
	Expense_frequency       *float64   `json:"expense_frequency"`
	Expense_frequency_uom   *string    `json:"expense_frequency_uom"`
	Expense_rate            *float64   `json:"expense_rate"`
	Expense_rate_qualifier  *string    `json:"expense_rate_qualifier"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ineligible_ind          *string    `json:"ineligible_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
