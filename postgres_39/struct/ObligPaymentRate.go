package dto

import (
	"time"
)

type Oblig_payment_rate struct {
	Jurisdiction       string     `json:"jurisdiction"`
	Pay_rate_type      string     `json:"pay_rate_type"`
	Payment_rate_id    string     `json:"payment_rate_id"`
	Active_ind         *string    `json:"active_ind"`
	Contract_id        *string    `json:"contract_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Payment_rate       *float64   `json:"payment_rate"`
	Payment_rate_ouom  *string    `json:"payment_rate_ouom"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Rate_percent       *float64   `json:"rate_percent"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
