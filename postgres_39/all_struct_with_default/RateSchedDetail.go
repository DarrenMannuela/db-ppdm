package dto

type Rate_sched_detail struct {
	Rate_schedule_id        string   `json:"rate_schedule_id" default:"source"`
	Rate_type               string   `json:"rate_type" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Currency_conversion     *float64 `json:"currency_conversion" default:""`
	Currency_ouom           *string  `json:"currency_ouom" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Nat_currency_conversion *float64 `json:"nat_currency_conversion" default:""`
	Nat_currency_uom        *string  `json:"nat_currency_uom" default:""`
	Period_type             *string  `json:"period_type" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Rate_condition          *string  `json:"rate_condition" default:""`
	Rate_cost               *float64 `json:"rate_cost" default:""`
	Rate_cost_uom           *string  `json:"rate_cost_uom" default:""`
	Rate_formula            *string  `json:"rate_formula" default:""`
	Rate_percent            *float64 `json:"rate_percent" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Taxable_ind             *string  `json:"taxable_ind" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
