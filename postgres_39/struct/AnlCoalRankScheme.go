package dto

import (
	"time"
)

type Anl_coal_rank_scheme struct {
	Coal_rank_scheme_id   string     `json:"coal_rank_scheme_id"`
	Active_ind            *string    `json:"active_ind"`
	Coal_rank_scheme_type *string    `json:"coal_rank_scheme_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Remark                *string    `json:"remark"`
	Scheme_name           *string    `json:"scheme_name"`
	Scheme_owner_ba_id    *string    `json:"scheme_owner_ba_id"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
