package dto

import (
	"time"
)

type Land_status struct {
	Land_right_subtype    string     `json:"land_right_subtype"`
	Land_right_id         string     `json:"land_right_id"`
	Status_type           string     `json:"status_type"`
	Land_right_status     string     `json:"land_right_status"`
	Status_seq_no         int        `json:"status_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Effective_term        *string    `json:"effective_term"`
	Effective_term_ouom   *string    `json:"effective_term_ouom"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Reason                *string    `json:"reason"`
	Remark                *string    `json:"remark"`
	Section_of_act        *string    `json:"section_of_act"`
	Section_of_act_date   *time.Time `json:"section_of_act_date"`
	Source                *string    `json:"source"`
	Undetermined_term_ind *string    `json:"undetermined_term_ind"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
