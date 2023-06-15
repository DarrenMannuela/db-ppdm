package dto

import (
	"time"
)

type Ba_xref struct {
	Business_associate_id string     `json:"business_associate_id"`
	New_ba_id             string     `json:"new_ba_id"`
	Ba_xref_seq_no        int        `json:"ba_xref_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Ba_xref_type          *string    `json:"ba_xref_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
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
