package dto

type Ppdm_map_detail struct {
	Map_id             string  `json:"map_id" default:"source"`
	Map_detail_obs_no  int     `json:"map_detail_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Column_name1       *string `json:"column_name_1" default:""`
	Column_name2       *string `json:"column_name_2" default:""`
	Create_method      *string `json:"create_method" default:""`
	Default_value      *string `json:"default_value" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Map_desc           *string `json:"map_desc" default:""`
	Map_owner_ba_id    *string `json:"map_owner_ba_id" default:""`
	Map_type           *string `json:"map_type" default:""`
	Map_version_num    *string `json:"map_version_num" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_ind      *string `json:"preferred_ind" default:""`
	Remark             *string `json:"remark" default:""`
	Ring_seq_no        *int    `json:"ring_seq_no" default:""`
	Schema_entity_id1  *string `json:"schema_entity_id_1" default:""`
	Schema_entity_id2  *string `json:"schema_entity_id_2" default:""`
	Source             *string `json:"source" default:""`
	Sw_application_id  *string `json:"sw_application_id" default:""`
	System_id1         *string `json:"system_id_1" default:""`
	System_id2         *string `json:"system_id_2" default:""`
	Table_name1        *string `json:"table_name_1" default:""`
	Table_name2        *string `json:"table_name_2" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
