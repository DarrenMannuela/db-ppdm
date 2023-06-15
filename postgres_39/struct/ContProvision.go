package dto

import (
	"time"
)

type Cont_provision struct {
	Contract_id         string     `json:"contract_id"`
	Provision_id        string     `json:"provision_id"`
	Active_ind          *string    `json:"active_ind"`
	Clause_heading      *string    `json:"clause_heading"`
	Clause_number       *string    `json:"clause_number"`
	Cont_provision_type *string    `json:"cont_provision_type"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Modified_ind        *string    `json:"modified_ind"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Provision_desc      *string    `json:"provision_desc"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Source_document_id  *string    `json:"source_document_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
