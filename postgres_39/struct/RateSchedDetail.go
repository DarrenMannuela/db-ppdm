package dto

import (
	"time"
)

type Rate_sched_detail struct {
	Rate_schedule_id        string     `json:"rate_schedule_id"`
	Rate_type               string     `json:"rate_type"`
	Active_ind              *string    `json:"active_ind"`
	Currency_conversion     *float64   `json:"currency_conversion"`
	Currency_ouom           *string    `json:"currency_ouom"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Nat_currency_conversion *float64   `json:"nat_currency_conversion"`
	Nat_currency_uom        *string    `json:"nat_currency_uom"`
	Period_type             *string    `json:"period_type"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Rate_condition          *string    `json:"rate_condition"`
	Rate_cost               *float64   `json:"rate_cost"`
	Rate_cost_uom           *string    `json:"rate_cost_uom"`
	Rate_formula            *string    `json:"rate_formula"`
	Rate_percent            *float64   `json:"rate_percent"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Taxable_ind             *string    `json:"taxable_ind"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
