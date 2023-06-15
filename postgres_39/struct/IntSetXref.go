package dto

import (
	"time"
)

type Int_set_xref struct {
	Interest_set_id       string     `json:"interest_set_id"`
	Interest_set_seq_no   int        `json:"interest_set_seq_no"`
	Interest_set_id_2     string     `json:"interest_set_id_2"`
	Interest_set_seq_no_2 int        `json:"interest_set_seq_no_2"`
	Int_set_xref_obs_no   int        `json:"int_set_xref_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Contract_id           *string    `json:"contract_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Int_set_xref_type     *string    `json:"int_set_xref_type"`
	Partner_ba_id         *string    `json:"partner_ba_id"`
	Partner_ba_id_2       *string    `json:"partner_ba_id_2"`
	Partner_obs_no        *int       `json:"partner_obs_no"`
	Partner_obs_no_2      *int       `json:"partner_obs_no_2"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Provision_id          *string    `json:"provision_id"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
