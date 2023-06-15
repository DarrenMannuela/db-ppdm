package dto

type Strat_hierarchy struct {
	Parent_strat_name_set_id   string  `json:"parent_strat_name_set_id" default:"source"`
	Parent_strat_unit_id       string  `json:"parent_strat_unit_id" default:"source"`
	Child_strat_name_set_id    string  `json:"child_strat_name_set_id" default:"source"`
	Child_strat_unit_id        string  `json:"child_strat_unit_id" default:"source"`
	Hierarchy_id               string  `json:"hierarchy_id" default:"source"`
	Source                     string  `json:"source" default:"source"`
	Active_ind                 *string `json:"active_ind" default:""`
	Area_id                    *string `json:"area_id" default:""`
	Area_type                  *string `json:"area_type" default:""`
	Child_interp_id2           *string `json:"child_interp_id_2" default:""`
	Child_strat_column_id      *string `json:"child_strat_column_id" default:""`
	Child_strat_column_source  *string `json:"child_strat_column_source" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Parent_interp_id           *string `json:"parent_interp_id" default:""`
	Parent_strat_column_id     *string `json:"parent_strat_column_id" default:""`
	Parent_strat_column_source *string `json:"parent_strat_column_source" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Preferred_hierarchy_ind    *string `json:"preferred_hierarchy_ind" default:""`
	Remark                     *string `json:"remark" default:""`
	Source_document_id         *string `json:"source_document_id" default:""`
	Strat_hierarchy_type       *string `json:"strat_hierarchy_type" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
