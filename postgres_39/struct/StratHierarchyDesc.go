package dto

import (
	"time"
)

type Strat_hierarchy_desc struct {
	Hierarchy_desc_id      string     `json:"hierarchy_desc_id"`
	Hierarchy_seq_no       int        `json:"hierarchy_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Parent_strat_unit_type *string    `json:"parent_strat_unit_type"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Source_document_id     *string    `json:"source_document_id"`
	Strat_hierarchy_type   *string    `json:"strat_hierarchy_type"`
	Strat_type             *string    `json:"strat_type"`
	Strat_unit_type        *string    `json:"strat_unit_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
