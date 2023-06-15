package dto

import (
	"time"
)

type Ppdm_map_rule struct {
	Map_id               string     `json:"map_id"`
	Map_detail_obs_no    int        `json:"map_detail_obs_no"`
	Rule_seq_no          int        `json:"rule_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Create_method        *string    `json:"create_method"`
	Date_format_desc     *time.Time `json:"date_format_desc"`
	Dep_column_name      *string    `json:"dep_column_name"`
	Dep_schema_entity_id *string    `json:"dep_schema_entity_id"`
	Dep_system_id        *string    `json:"dep_system_id"`
	Dep_table_name       *string    `json:"dep_table_name"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Map_rule_type        *string    `json:"map_rule_type"`
	Max_value            *float64   `json:"max_value"`
	Max_value_ouom       *string    `json:"max_value_ouom"`
	Max_value_uom        *string    `json:"max_value_uom"`
	Min_value            *float64   `json:"min_value"`
	Min_value_ouom       *string    `json:"min_value_ouom"`
	Min_value_uom        *string    `json:"min_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Procedure_id         *string    `json:"procedure_id"`
	Procedure_system_id  *string    `json:"procedure_system_id"`
	Remark               *string    `json:"remark"`
	Ring_seq_no          *int       `json:"ring_seq_no"`
	Rule_desc            *string    `json:"rule_desc"`
	Rule_owner_ba_id     *string    `json:"rule_owner_ba_id"`
	Rule_version_num     *string    `json:"rule_version_num"`
	Source               *string    `json:"source"`
	Sw_application_id    *string    `json:"sw_application_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
