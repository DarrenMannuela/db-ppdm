package dto

import (
	"time"
)

type Oblig_payment_instr struct {
	Payee_ba_id            string     `json:"payee_ba_id"`
	Payment_instruction_id string     `json:"payment_instruction_id"`
	Aba_number             *string    `json:"aba_number"`
	Active_ind             *string    `json:"active_ind"`
	Bank_address_obs_no    *int       `json:"bank_address_obs_no"`
	Bank_address_source    *string    `json:"bank_address_source"`
	Bank_ba_id             *string    `json:"bank_ba_id"`
	Bank_fee               *float64   `json:"bank_fee"`
	Currency_conversion    *float64   `json:"currency_conversion"`
	Currency_ouom          *string    `json:"currency_ouom"`
	Depository_num         *string    `json:"depository_num"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Invalid_date           *time.Time `json:"invalid_date"`
	Pay_method             *string    `json:"pay_method"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Suspend_payment_ind    *string    `json:"suspend_payment_ind"`
	Suspend_payment_reason *string    `json:"suspend_payment_reason"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
