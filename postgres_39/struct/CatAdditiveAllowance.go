package dto

import (
	"time"
)

type Cat_additive_allowance struct {
	Allowance_id          string     `json:"allowance_id"`
	Active_ind            *string    `json:"active_ind"`
	Additive_group_id     *string    `json:"additive_group_id"`
	Allowed_ind           *string    `json:"allowed_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Catalogue_additive_id *string    `json:"catalogue_additive_id"`
	Disallowed_ind        *string    `json:"disallowed_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Source_document_id    *string    `json:"source_document_id"`
	Substance_id          *string    `json:"substance_id"`
	Use_condition         *string    `json:"use_condition"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
