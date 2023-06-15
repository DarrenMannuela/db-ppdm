package dto

import (
	"time"
)

type Int_set_partner struct {
	Interest_set_id        string     `json:"interest_set_id"`
	Interest_set_seq_no    int        `json:"interest_set_seq_no"`
	Partner_ba_id          string     `json:"partner_ba_id"`
	Partner_obs_no         int        `json:"partner_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Address_obs_no         *int       `json:"address_obs_no"`
	Address_source         *string    `json:"address_source"`
	Breach_desc            *string    `json:"breach_desc"`
	Breach_ind             *string    `json:"breach_ind"`
	Change_reason          *string    `json:"change_reason"`
	Confidential_ind       *string    `json:"confidential_ind"`
	Contract_id            *string    `json:"contract_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Escrow_desc            *string    `json:"escrow_desc"`
	Escrow_ind             *string    `json:"escrow_ind"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Gross_fract_interest   *string    `json:"gross_fract_interest"`
	Gross_percent_interest *float64   `json:"gross_percent_interest"`
	Inactive_date          *time.Time `json:"inactive_date"`
	Instrument_id          *string    `json:"instrument_id"`
	Interest_set_role      *string    `json:"interest_set_role"`
	Net_fract_interest     *string    `json:"net_fract_interest"`
	Net_percent_interest   *float64   `json:"net_percent_interest"`
	Penalty_ind            *string    `json:"penalty_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Provision_id           *string    `json:"provision_id"`
	Remark                 *string    `json:"remark"`
	Right_to_earn_desc     *string    `json:"right_to_earn_desc"`
	Right_to_earn_ind      *string    `json:"right_to_earn_ind"`
	Source                 *string    `json:"source"`
	Title_own_type         *string    `json:"title_own_type"`
	Unit_operated_ind      *string    `json:"unit_operated_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
