package dto

import (
	"time"
)

type Land_size struct {
	Land_right_subtype    string     `json:"land_right_subtype"`
	Land_right_id         string     `json:"land_right_id"`
	Size_type             string     `json:"size_type"`
	Size_seq_no           int        `json:"size_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Gross_size            *float64   `json:"gross_size"`
	Interest_set_id       *string    `json:"interest_set_id"`
	Interest_set_seq_no   *int       `json:"interest_set_seq_no"`
	Net_size              *float64   `json:"net_size"`
	Partner_obs_no        *int       `json:"partner_obs_no"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Size_ouom             *string    `json:"size_ouom"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
