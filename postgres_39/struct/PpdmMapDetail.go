package dto

import (
	"time"
)

type Ppdm_map_detail struct {
	Map_id             string     `json:"map_id"`
	Map_detail_obs_no  int        `json:"map_detail_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Column_name1       *string    `json:"column_name_1"`
	Column_name2       *string    `json:"column_name_2"`
	Create_method      *string    `json:"create_method"`
	Default_value      *string    `json:"default_value"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Map_desc           *string    `json:"map_desc"`
	Map_owner_ba_id    *string    `json:"map_owner_ba_id"`
	Map_type           *string    `json:"map_type"`
	Map_version_num    *string    `json:"map_version_num"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_ind      *string    `json:"preferred_ind"`
	Remark             *string    `json:"remark"`
	Ring_seq_no        *int       `json:"ring_seq_no"`
	Schema_entity_id1  *string    `json:"schema_entity_id_1"`
	Schema_entity_id2  *string    `json:"schema_entity_id_2"`
	Source             *string    `json:"source"`
	Sw_application_id  *string    `json:"sw_application_id"`
	System_id1         *string    `json:"system_id_1"`
	System_id2         *string    `json:"system_id_2"`
	Table_name1        *string    `json:"table_name_1"`
	Table_name2        *string    `json:"table_name_2"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
