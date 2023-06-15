package dto

import (
	"time"
)

type Int_set_status struct {
	Interest_set_id       string     `json:"interest_set_id"`
	Interest_set_seq_no   int        `json:"interest_set_seq_no"`
	Status_obs_no         int        `json:"status_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Effective_term        *string    `json:"effective_term"`
	Effective_term_ouom   *string    `json:"effective_term_ouom"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Reason                *string    `json:"reason"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Status                *string    `json:"status"`
	Status_type           *string    `json:"status_type"`
	Undetermined_term_ind *string    `json:"undetermined_term_ind"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
