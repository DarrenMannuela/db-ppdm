package dto

import (
	"time"
)

type Fin_cost_summary struct {
	Cost_id              string     `json:"cost_id"`
	Active_ind           *string    `json:"active_ind"`
	Ami_ind              *string    `json:"ami_ind"`
	Confidential_ind     *string    `json:"confidential_ind"`
	Cost_amount          *float64   `json:"cost_amount"`
	Cost_per_size        *float64   `json:"cost_per_size"`
	Cost_per_size_ouom   *string    `json:"cost_per_size_ouom"`
	Cost_type            *string    `json:"cost_type"`
	Currency_conversion  *float64   `json:"currency_conversion"`
	Currency_ouom        *string    `json:"currency_ouom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Finance_component_id *string    `json:"finance_component_id"`
	Finance_id           *string    `json:"finance_id"`
	Paid_date            *time.Time `json:"paid_date"`
	Parent_cost_id       *string    `json:"parent_cost_id"`
	Percent_of_parent    *float64   `json:"percent_of_parent"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Report_cost_ind      *string    `json:"report_cost_ind"`
	Source               *string    `json:"source"`
	Submit_date          *time.Time `json:"submit_date"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
