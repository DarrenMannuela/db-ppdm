package dto

import (
	"time"
)

type Finance struct {
	Finance_id          string     `json:"finance_id"`
	Active_ind          *string    `json:"active_ind"`
	Actual_cost         *float64   `json:"actual_cost"`
	Authorized_by_ba_id *string    `json:"authorized_by_ba_id"`
	Budget_cost         *float64   `json:"budget_cost"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Finance_type        *string    `json:"finance_type"`
	Fin_status          *string    `json:"fin_status"`
	Issue_date          *time.Time `json:"issue_date"`
	Limit_amount        *float64   `json:"limit_amount"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Reference_number    *string    `json:"reference_number"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Tax_credit_code     *string    `json:"tax_credit_code"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
