package dto

import (
	"time"
)

type Strat_hierarchy struct {
	Parent_strat_name_set_id   string     `json:"parent_strat_name_set_id"`
	Parent_strat_unit_id       string     `json:"parent_strat_unit_id"`
	Child_strat_name_set_id    string     `json:"child_strat_name_set_id"`
	Child_strat_unit_id        string     `json:"child_strat_unit_id"`
	Hierarchy_id               string     `json:"hierarchy_id"`
	Source                     string     `json:"source"`
	Active_ind                 *string    `json:"active_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_type                  *string    `json:"area_type"`
	Child_interp_id2           *string    `json:"child_interp_id_2"`
	Child_strat_column_id      *string    `json:"child_strat_column_id"`
	Child_strat_column_source  *string    `json:"child_strat_column_source"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Parent_interp_id           *string    `json:"parent_interp_id"`
	Parent_strat_column_id     *string    `json:"parent_strat_column_id"`
	Parent_strat_column_source *string    `json:"parent_strat_column_source"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_hierarchy_ind    *string    `json:"preferred_hierarchy_ind"`
	Remark                     *string    `json:"remark"`
	Source_document_id         *string    `json:"source_document_id"`
	Strat_hierarchy_type       *string    `json:"strat_hierarchy_type"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
