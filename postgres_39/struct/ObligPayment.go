package dto

import (
	"time"
)

type Oblig_payment struct {
	Obligation_id             string     `json:"obligation_id"`
	Obligation_seq_no         int        `json:"obligation_seq_no"`
	Payee_payor_ba_id         string     `json:"payee_payor_ba_id"`
	Payee_payor               string     `json:"payee_payor"`
	Active_ind                *string    `json:"active_ind"`
	Cheque_remark             *string    `json:"cheque_remark"`
	Contract_id               *string    `json:"contract_id"`
	Currency_conversion       *float64   `json:"currency_conversion"`
	Currency_ouom             *string    `json:"currency_ouom"`
	Currency_uom              *string    `json:"currency_uom"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Gross_cost                *float64   `json:"gross_cost"`
	Invoice_amount            *float64   `json:"invoice_amount"`
	Invoice_num               *string    `json:"invoice_num"`
	Land_rental_type          *string    `json:"land_rental_type"`
	Net_cost                  *float64   `json:"net_cost"`
	Payment_due_date          *time.Time `json:"payment_due_date"`
	Payment_ind               *string    `json:"payment_ind"`
	Payment_instruction_id    *string    `json:"payment_instruction_id"`
	Payment_type              *string    `json:"payment_type"`
	Pay_method                *string    `json:"pay_method"`
	Percent_of_gross          *float64   `json:"percent_of_gross"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Provision_id              *string    `json:"provision_id"`
	Rate_schedule_id          *string    `json:"rate_schedule_id"`
	Rate_type                 *string    `json:"rate_type"`
	Receipt_ind               *string    `json:"receipt_ind"`
	Related_obligation_id     *string    `json:"related_obligation_id"`
	Related_obligation_seq_no *int       `json:"related_obligation_seq_no"`
	Related_oblig_ba_id       *string    `json:"related_oblig_ba_id"`
	Related_oblig_payee_payor *string    `json:"related_oblig_payee_payor"`
	Remark                    *string    `json:"remark"`
	Royalty_type              *string    `json:"royalty_type"`
	Source                    *string    `json:"source"`
	Stop_payment_ind          *string    `json:"stop_payment_ind"`
	Substance_id              *string    `json:"substance_id"`
	Substance_seq_no          *int       `json:"substance_seq_no"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
