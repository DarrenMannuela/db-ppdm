package dto

import (
	"time"
)

type Ppdm_schema_group struct {
	Group_system_id        string     `json:"group_system_id"`
	Group_schema_entity_id string     `json:"group_schema_entity_id"`
	Comp_schema_entity_id  string     `json:"comp_schema_entity_id"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Entity_seq_no          *int       `json:"entity_seq_no"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Extension_ind          *string    `json:"extension_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Schema_group_type      *string    `json:"schema_group_type"`
	Source                 *string    `json:"source"`
	Surrogate_ind          *string    `json:"surrogate_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
