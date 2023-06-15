package dto

import (
	"time"
)

type Obligation struct {
	Obligation_id             string     `json:"obligation_id"`
	Obligation_seq_no         int        `json:"obligation_seq_no"`
	Active_ind                *string    `json:"active_ind"`
	Calculation_method        *string    `json:"calculation_method"`
	Convertible_ind           *string    `json:"convertible_ind"`
	Critical_date             *time.Time `json:"critical_date"`
	Currency_conversion       *float64   `json:"currency_conversion"`
	Currency_ouom             *string    `json:"currency_ouom"`
	Description               *string    `json:"description"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Fulfilled_date            *time.Time `json:"fulfilled_date"`
	Fulfilled_ind             *string    `json:"fulfilled_ind"`
	Gross_obligation_cost     *float64   `json:"gross_obligation_cost"`
	Instrument_id             *string    `json:"instrument_id"`
	Liability_release_date    *time.Time `json:"liability_release_date"`
	Net_obligation_cost       *float64   `json:"net_obligation_cost"`
	Notice_period_length      *float64   `json:"notice_period_length"`
	Notice_period_ouom        *string    `json:"notice_period_ouom"`
	Obligation_category       *string    `json:"obligation_category"`
	Obligation_duration       *float64   `json:"obligation_duration"`
	Obligation_duration_ouom  *string    `json:"obligation_duration_ouom"`
	Obligation_frequency      *string    `json:"obligation_frequency"`
	Obligation_type           *string    `json:"obligation_type"`
	Payment_ind               *string    `json:"payment_ind"`
	Payment_responsibility    *string    `json:"payment_responsibility"`
	Percentage                *float64   `json:"percentage"`
	Potential_obligation_desc *string    `json:"potential_obligation_desc"`
	Potential_obligation_ind  *string    `json:"potential_obligation_ind"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Prepaid_ind               *string    `json:"prepaid_ind"`
	Remark                    *string    `json:"remark"`
	Resp_party_ba_id          *string    `json:"resp_party_ba_id"`
	Review_frequency          *string    `json:"review_frequency"`
	Royalty_owner_ba_id       *string    `json:"royalty_owner_ba_id"`
	Royalty_payor_ba_id       *string    `json:"royalty_payor_ba_id"`
	Royalty_type              *string    `json:"royalty_type"`
	Source                    *string    `json:"source"`
	Start_date                *time.Time `json:"start_date"`
	Substance_id              *string    `json:"substance_id"`
	Trigger_method            *string    `json:"trigger_method"`
	Work_obligation_desc      *string    `json:"work_obligation_desc"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
