package dto

import (
	"time"
)

type Oblig_pay_detail struct {
	Obligation_id       string     `json:"obligation_id"`
	Obligation_seq_no   int        `json:"obligation_seq_no"`
	Payee_payor_ba_id   string     `json:"payee_payor_ba_id"`
	Payee_payor         string     `json:"payee_payor"`
	Detail_id           string     `json:"detail_id"`
	Active_ind          *string    `json:"active_ind"`
	Cheque_number       *string    `json:"cheque_number"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Deduction_ind       *string    `json:"deduction_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Payment_amount      *float64   `json:"payment_amount"`
	Payment_date        *time.Time `json:"payment_date"`
	Pay_detail_type     *string    `json:"pay_detail_type"`
	Percent_of_payment  *float64   `json:"percent_of_payment"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Rate                *float64   `json:"rate"`
	Rate_ouom           *string    `json:"rate_ouom"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
