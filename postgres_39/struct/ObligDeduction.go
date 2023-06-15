package dto

import (
	"time"
)

type Oblig_deduction struct {
	Obligation_id         string     `json:"obligation_id"`
	Obligation_seq_no     int        `json:"obligation_seq_no"`
	Deduction_id          string     `json:"deduction_id"`
	Active_ind            *string    `json:"active_ind"`
	Allow_deduction_id    *string    `json:"allow_deduction_id"`
	Currency_conversion   *float64   `json:"currency_conversion"`
	Currency_ouom         *string    `json:"currency_ouom"`
	Deduction_amount      *float64   `json:"deduction_amount"`
	Deduction_percent     *float64   `json:"deduction_percent"`
	Deduct_type           *string    `json:"deduct_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Max_deduction_allowed *float64   `json:"max_deduction_allowed"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Royalty_amount        *float64   `json:"royalty_amount"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
