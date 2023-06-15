package dto

import (
	"time"
)

type Fossil_taxon_hier struct {
	Parent_taxon_leaf_id string     `json:"parent_taxon_leaf_id"`
	Child_taxon_leaf_id  string     `json:"child_taxon_leaf_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Taxon_name           *string    `json:"taxon_name"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
