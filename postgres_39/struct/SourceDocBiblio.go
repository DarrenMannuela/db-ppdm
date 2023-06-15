package dto

import (
	"time"
)

type Source_doc_biblio struct {
	Source_document_id  string     `json:"source_document_id"`
	Biblio_obs_no       int        `json:"biblio_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Document_name       *string    `json:"document_name"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Referenced_document *string    `json:"referenced_document"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
