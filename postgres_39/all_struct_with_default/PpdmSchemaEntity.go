package dto

type Ppdm_schema_entity struct {
	System_id          string  `json:"system_id" default:"source"`
	Schema_entity_id   string  `json:"schema_entity_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Data_type          *string `json:"data_type" default:""`
	Default_uom_id     *string `json:"default_uom_id" default:""`
	Delimiter          *string `json:"delimiter" default:""`
	Domain             *string `json:"domain" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Element_type       *string `json:"element_type" default:""`
	Entity_comment     *string `json:"entity_comment" default:""`
	Entity_precision   *int    `json:"entity_precision" default:""`
	Entity_scale       *int    `json:"entity_scale" default:""`
	Entity_seq_no      *int    `json:"entity_seq_no" default:""`
	Entity_size        *int    `json:"entity_size" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Extension_ind      *string `json:"extension_ind" default:""`
	Last_system_key    *string `json:"last_system_key" default:""`
	Nullable_ind       *string `json:"nullable_ind" default:""`
	Position_end       *int    `json:"position_end" default:""`
	Position_start     *int    `json:"position_start" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_name     *string `json:"preferred_name" default:""`
	Reference_num      *string `json:"reference_num" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Surrogate_ind      *string `json:"surrogate_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
