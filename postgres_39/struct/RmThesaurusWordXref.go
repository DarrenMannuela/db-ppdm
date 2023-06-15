package dto

import (
	"time"
)

type Rm_thesaurus_word_xref struct {
	Thesaurus_id1            string     `json:"thesaurus_id_1"`
	Thesaurus_word_id1       string     `json:"thesaurus_word_id_1"`
	Thesaurus_id2            string     `json:"thesaurus_id_2"`
	Thesaurus_word_id2       string     `json:"thesaurus_word_id_2"`
	Xref_obs_no              int        `json:"xref_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Relationship_description *string    `json:"relationship_description"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	Thesaurus_xref_type      *string    `json:"thesaurus_xref_type"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
