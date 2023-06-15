package dto

type Work_order_ba struct {
	Work_order_id         string   `json:"work_order_id" default:"source"`
	Business_associate_id string   `json:"business_associate_id" default:"source"`
	Ba_obs_no             int      `json:"ba_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Ba_role               *string  `json:"ba_role" default:""`
	Currency_conversion   *float64 `json:"currency_conversion" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Payment_amount        *float64 `json:"payment_amount" default:""`
	Payment_amount_ouom   *string  `json:"payment_amount_ouom" default:""`
	Payment_percent       *float64 `json:"payment_percent" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
