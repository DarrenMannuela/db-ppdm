package dto

import (
	"time"
)

type Ppdm_schema_entity struct {
	System_id          string     `json:"system_id"`
	Schema_entity_id   string     `json:"schema_entity_id"`
	Active_ind         *string    `json:"active_ind"`
	Data_type          *string    `json:"data_type"`
	Default_uom_id     *string    `json:"default_uom_id"`
	Delimiter          *string    `json:"delimiter"`
	Domain             *string    `json:"domain"`
	Effective_date     *time.Time `json:"effective_date"`
	Element_type       *string    `json:"element_type"`
	Entity_comment     *string    `json:"entity_comment"`
	Entity_precision   *int       `json:"entity_precision"`
	Entity_scale       *int       `json:"entity_scale"`
	Entity_seq_no      *int       `json:"entity_seq_no"`
	Entity_size        *int       `json:"entity_size"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Extension_ind      *string    `json:"extension_ind"`
	Last_system_key    *string    `json:"last_system_key"`
	Nullable_ind       *string    `json:"nullable_ind"`
	Position_end       *int       `json:"position_end"`
	Position_start     *int       `json:"position_start"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_name     *string    `json:"preferred_name"`
	Reference_num      *string    `json:"reference_num"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Surrogate_ind      *string    `json:"surrogate_ind"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
