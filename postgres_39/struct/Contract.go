package dto

import (
	"time"
)

type Contract struct {
	Contract_id                string     `json:"contract_id"`
	Active_ind                 *string    `json:"active_ind"`
	Ami_aoe_ind                *string    `json:"ami_aoe_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_type                  *string    `json:"area_type"`
	Assignment_proc_ind        *string    `json:"assignment_proc_ind"`
	Casing_point_election      *string    `json:"casing_point_election"`
	Closing_date               *time.Time `json:"closing_date"`
	Confidential_ind           *string    `json:"confidential_ind"`
	Contract_name              *string    `json:"contract_name"`
	Contract_num               *string    `json:"contract_num"`
	Currency_conversion        *float64   `json:"currency_conversion"`
	Currency_ouom              *string    `json:"currency_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Extension_condition        *string    `json:"extension_condition"`
	Governing_contract_ind     *string    `json:"governing_contract_ind"`
	Liability_period           *float64   `json:"liability_period"`
	Liability_period_ouom      *string    `json:"liability_period_ouom"`
	Liability_release_date     *time.Time `json:"liability_release_date"`
	Nat_currency_conversion    *float64   `json:"nat_currency_conversion"`
	Nat_currency_uom           *string    `json:"nat_currency_uom"`
	Operating_proc_ind         *string    `json:"operating_proc_ind"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Primary_term               *float64   `json:"primary_term"`
	Primary_term_ouom          *string    `json:"primary_term_ouom"`
	Remark                     *string    `json:"remark"`
	Rofr_ind                   *string    `json:"rofr_ind"`
	Secondary_term             *float64   `json:"secondary_term"`
	Secondary_term_ouom        *string    `json:"secondary_term_ouom"`
	Source                     *string    `json:"source"`
	Source_document            *string    `json:"source_document"`
	Surrender_notice_term      *float64   `json:"surrender_notice_term"`
	Surrender_notice_term_ouom *string    `json:"surrender_notice_term_ouom"`
	Termination_date           *time.Time `json:"termination_date"`
	Voting_proc_ind            *string    `json:"voting_proc_ind"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
