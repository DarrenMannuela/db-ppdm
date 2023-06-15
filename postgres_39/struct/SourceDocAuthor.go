package dto

import (
	"time"
)

type Source_doc_author struct {
	Source_document_id string     `json:"source_document_id"`
	Author_id          string     `json:"author_id"`
	Active_ind         *string    `json:"active_ind"`
	Author_ba_id       *string    `json:"author_ba_id"`
	Author_first_name  *string    `json:"author_first_name"`
	Author_initial     *string    `json:"author_initial"`
	Author_last_name   *string    `json:"author_last_name"`
	Author_seq_no      *int       `json:"author_seq_no"`
	Author_type        *string    `json:"author_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
