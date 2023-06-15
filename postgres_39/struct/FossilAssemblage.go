package dto

import (
	"time"
)

type Fossil_assemblage struct {
	Strat_name_set_id  string     `json:"strat_name_set_id"`
	Strat_unit_id      string     `json:"strat_unit_id"`
	Taxon_leaf_id      string     `json:"taxon_leaf_id"`
	Interp_id          string     `json:"interp_id"`
	Active_ind         *string    `json:"active_ind"`
	Assemblage_type    *string    `json:"assemblage_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Oldest_ind         *string    `json:"oldest_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_marker_ind *string    `json:"primary_marker_ind"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
