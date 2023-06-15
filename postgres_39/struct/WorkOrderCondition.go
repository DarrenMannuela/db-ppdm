package dto

import (
	"time"
)

type Work_order_condition struct {
	Work_order_id         string     `json:"work_order_id"`
	Condition_obs_no      int        `json:"condition_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Ba_role               *string    `json:"ba_role"`
	Business_associate_id *string    `json:"business_associate_id"`
	Condition_desc        *string    `json:"condition_desc"`
	Condition_type        *string    `json:"condition_type"`
	Currency_conversion   *float64   `json:"currency_conversion"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Payment_amount        *float64   `json:"payment_amount"`
	Payment_amount_ouom   *string    `json:"payment_amount_ouom"`
	Payment_percent       *float64   `json:"payment_percent"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Work_order_type       *string    `json:"work_order_type"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
