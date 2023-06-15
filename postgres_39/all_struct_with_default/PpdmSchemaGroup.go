package dto

type Ppdm_schema_group struct {
	Group_system_id        string  `json:"group_system_id" default:"source"`
	Group_schema_entity_id string  `json:"group_schema_entity_id" default:"source"`
	Comp_schema_entity_id  string  `json:"comp_schema_entity_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Entity_seq_no          *int    `json:"entity_seq_no" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Extension_ind          *string `json:"extension_ind" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Schema_group_type      *string `json:"schema_group_type" default:""`
	Source                 *string `json:"source" default:""`
	Surrogate_ind          *string `json:"surrogate_ind" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
